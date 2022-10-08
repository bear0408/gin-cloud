module ginweb

go 1.19

require (
	github.com/asim/go-micro/v3 v3.7.1
	github.com/gin-gonic/gin v1.8.1
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/google/uuid v1.3.0
	github.com/hashicorp/consul/api v1.12.0
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/polarismesh/polaris-go v1.2.0-beta.3
	github.com/spf13/viper v1.13.0
	go.uber.org/zap v1.23.0
	gorm.io/driver/mysql v1.3.6
	gorm.io/gorm v1.23.9
)

require google.golang.org/protobuf v1.28.1

require (
	github.com/armon/go-metrics v0.3.10 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dlclark/regexp2 v1.7.0 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.2.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/serf v0.9.7 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.20.2 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.13.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/afero v1.8.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/net v0.0.0-20220920203100-d0c6ba3f52d9 // indirect
	golang.org/x/sync v0.0.0-20220601150217-0de741cfad7f // indirect
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220920201722-2b89144ce006 // indirect
	google.golang.org/grpc v1.49.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

//require (
//	github.com/fsnotify/fsnotify v1.5.4 // indirect
//	github.com/gin-contrib/sse v0.1.0 // indirect
//	github.com/go-playground/locales v0.14.0 // indirect
//	github.com/go-playground/universal-translator v0.18.0 // indirect
//	github.com/go-playground/validator/v10 v10.11.1 // indirect
//	github.com/goccy/go-json v0.9.11 // indirect
//	github.com/hashicorp/hcl v1.0.0 // indirect
//	github.com/json-iterator/go v1.1.12 // indirect
//	github.com/leodido/go-urn v1.2.1 // indirect
//	github.com/magiconair/properties v1.8.6 // indirect
//	github.com/mattn/go-isatty v0.0.16 // indirect
//	github.com/mitchellh/mapstructure v1.5.0 // indirect
//	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
//	github.com/modern-go/reflect2 v1.0.2 // indirect
//	github.com/pelletier/go-toml v1.9.5 // indirect
//	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
//	github.com/spf13/afero v1.8.2 // indirect
//	github.com/spf13/cast v1.5.0 // indirect
//	github.com/spf13/jwalterweatherman v1.1.0 // indirect
//	github.com/spf13/pflag v1.0.5 // indirect
//	github.com/subosito/gotenv v1.4.1 // indirect
//	github.com/ugorji/go/codec v1.2.7 // indirect
//	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
//	golang.org/x/net v0.0.0-20220909164309-bea034e7d591 // indirect
//	golang.org/x/sys v0.0.0-20220915200043-7b5979e65e41 // indirect
//	golang.org/x/text v0.3.7 // indirect
//	google.golang.org/protobuf v1.28.1 // indirect
//	gopkg.in/ini.v1 v1.67.0 // indirect
//	gopkg.in/yaml.v2 v2.4.0 // indirect
//	gopkg.in/yaml.v3 v3.0.1 // indirect
//)
