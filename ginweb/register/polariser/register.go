/**
 * Tencent is pleased to support the open source community by making polaris-go available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package registry

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/polarismesh/polaris-go"
)

type RegistryInfo struct {
	Namespace string
	Service   string
	Host      string
	Token     string
	Port      int
	TtlSec    int
	cancel    context.CancelFunc
}

func (r *RegistryInfo) Key() string {
	return fmt.Sprintf("%s@%s@%s@%d", r.Namespace, r.Service, r.Host, r.Port)
}

// PolarisProvider is an example of provider
type RegistryClient struct {
	provider      polaris.ProviderAPI
	registryInfos map[string]*RegistryInfo
}

type option func(r *RegistryClient)

func NewRegistryClient() (*RegistryClient, error) {
	p, err := polaris.NewProviderAPI()
	if err != nil {
		return nil, err
	}
	reg := &RegistryClient{
		provider: p,
	}
	return reg, nil
}

func (c *RegistryClient) Register(info *RegistryInfo) {
	log.Printf("start to invoke register operation")
	registerRequest := &polaris.InstanceRegisterRequest{}
	registerRequest.Service = info.Service
	registerRequest.Namespace = info.Namespace
	registerRequest.Host = info.Host
	registerRequest.Port = info.Port
	registerRequest.ServiceToken = info.Token
	registerRequest.SetTTL(info.TtlSec)
	resp, err := c.provider.Register(registerRequest)
	if err != nil {
		log.Fatalf("fail to register instance, err is %v", err)
	}
	log.Printf("register response: instanceId %s", resp.InstanceID)
	c.doHeartbeat(info)
}

func (c *RegistryClient) doHeartbeat(info *RegistryInfo) {
	c.registryInfos[info.Key()] = info
	ctx, cancel := context.WithCancel(context.Background())

	info.cancel = cancel

	go func(ctx context.Context, info *RegistryInfo) {
		ticker := time.NewTicker(time.Duration(info.TtlSec) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				heartbeatRequest := &polaris.InstanceHeartbeatRequest{}
				heartbeatRequest.Service = info.Service
				heartbeatRequest.Namespace = info.Namespace
				heartbeatRequest.Host = info.Host
				heartbeatRequest.Port = info.Port

				if err := c.provider.Heartbeat(heartbeatRequest); err != nil {
					// TODO log.error
				}
			case <-ctx.Done():
				return
			}
		}
	}(ctx, info)
}

func (c *RegistryClient) UnRegister(info *RegistryInfo) {
	log.Printf("start to invoke deregister operation")

	save, ok := c.registryInfos[info.Key()]
	if ok && save.cancel != nil {
		save.cancel()
	}

	deregisterRequest := &polaris.InstanceDeRegisterRequest{}
	deregisterRequest.Service = info.Service
	deregisterRequest.Namespace = info.Namespace
	deregisterRequest.Host = info.Host
	deregisterRequest.Port = info.Port
	deregisterRequest.ServiceToken = info.Token
	if err := c.provider.Deregister(deregisterRequest); err != nil {
		log.Fatalf("fail to deregister instance, err is %v", err)
	}
	log.Printf("deregister successfully.")
}
