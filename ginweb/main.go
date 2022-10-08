package main

import (
	"flag"
	"fmt"
	"ginweb/common"
	"ginweb/conf"
	"ginweb/initcon"
	_ "ginweb/initcon"
	"ginweb/middleware"
	"github.com/polarismesh/polaris-go"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	//	"github.com/polarismesh/grpc-go-polaris/examples/common/pb"

	//"ginweb/register/consul"
	//_ "ginweb/register/polariser"
	"ginweb/routes"
	//_ "github.com/asim/go-micro/plugins/server/grpc/v3"
	//_ "github.com/asim/go-micro/v3"        // "github.com/micro/micro/v3/service"
	//_ "github.com/asim/go-micro/v3/logger" // "github.com/micro/micro/v3/service/logger"
	//_ "github.com/asim/go-micro/v3/registry"
	"github.com/gin-gonic/gin"
	//_ "github.com/polarismesh/grpc-go-polaris/examples/common/pb"
	//_ "google.golang.org/grpc"
)

//const (
//	ServerName = "ginCloud.srv"
//	ConsulAddr = "192.168.233.130:8500"
//)

const (
	listenPort = 8088
)

var (
	namespace string
	service   string
	token     string
)

func initArgs() {
	flag.StringVar(&namespace, "namespace", "default", "namespace")
	flag.StringVar(&service, "service", "DiscoverEchoServer2", "service")
	// 当北极星开启鉴权时，需要配置此参数完成相关的权限检查
	flag.StringVar(&token, "token", "", "token")
}

// PolarisProvider is an example of provider
type PolarisProvider struct {
	provider   polaris.ProviderAPI
	namespace  string
	service    string
	host       string
	port       int
	isShutdown bool
}

// Run starts the provider
func (svr *PolarisProvider) Run() {
	//tmpHost, err := getLocalHost(svr.provider.SDKContext().GetConfig().GetGlobal().GetServerConnector().GetAddresses()[0])
	//if err != nil {
	//	panic(fmt.Errorf("error occur while fetching localhost: %v", err))
	//}

	svr.host = "10.107.9.140" //tmpHost
	svr.runWebServer()
	svr.registerService()
}

func (svr *PolarisProvider) runWebServer() {
	http.HandleFunc("/echo", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(fmt.Sprintf("Hello, I'm DiscoverEchoServer Provider, My host : %s:%d", svr.host, svr.port)))
	})

	ln, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		log.Fatalf("[ERROR]fail to listen tcp, err is %v", err)
	}

	svr.port = ln.Addr().(*net.TCPAddr).Port

	go func() {
		if err := http.Serve(ln, nil); err != nil {
			svr.isShutdown = false
			log.Fatalf("[ERROR]fail to run webServer, err is %v", err)
		}
	}()
}

func (svr *PolarisProvider) registerService() {
	log.Printf("start to invoke register operation")
	registerRequest := &polaris.InstanceRegisterRequest{}
	registerRequest.Service = service
	registerRequest.Namespace = namespace
	registerRequest.Host = svr.host
	registerRequest.Port = svr.port
	registerRequest.ServiceToken = token
	registerRequest.SetTTL(10)
	resp, err := svr.provider.RegisterInstance(registerRequest)
	if err != nil {
		log.Fatalf("fail to register instance, err is %v", err)
	}
	log.Printf("register response: instanceId %s", resp.InstanceID)
}

func (svr *PolarisProvider) deregisterService() {
	log.Printf("start to invoke deregister operation")
	deregisterRequest := &polaris.InstanceDeRegisterRequest{}
	deregisterRequest.Service = service
	deregisterRequest.Namespace = namespace
	deregisterRequest.Host = svr.host
	deregisterRequest.Port = svr.port
	deregisterRequest.ServiceToken = token
	if err := svr.provider.Deregister(deregisterRequest); err != nil {
		log.Fatalf("fail to deregister instance, err is %v", err)
	}
	log.Printf("deregister successfully.")
}

func getLocalHost(serverAddr string) (string, error) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return "", err
	}
	localAddr := conn.LocalAddr().String()
	colonIdx := strings.LastIndex(localAddr, ":")
	if colonIdx > 0 {
		return localAddr[:colonIdx], nil
	}
	return localAddr, nil
}

func (svr *PolarisProvider) runMainLoop() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, []os.Signal{
		syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGSEGV,
	}...)

	for s := range ch {
		log.Printf("catch signal(%+v), stop servers", s)
		svr.isShutdown = true
		svr.deregisterService()
		return
	}
}

func main() {

	//初始化配置
	config := conf.InitConf()

	//初始化mysql
	sqlConn := config.SqlConn
	//fmt.Println("mysql conn:",sqlConn)
	initcon.InitDB(sqlConn) //gorm

	//初始化redis
	//fmt.Println("redis db num:",config.RedisConf.Db)
	initcon.InitRedis(config.RedisConf)

	//初始化日志
	common.InitLogger()

	//_ = common.InitCasbinEnforcer()

	//初始化路由
	//route := gin.Default()
	route := gin.New()
	route.Use(middleware.GinLogger(), middleware.GinRecovery(false))
	//使用中间件解决跨域问题
	route.Use(middleware.Cors())

	routes.InitRouter(route)

	//项目启动
	//gin.SetMode(gin.ReleaseMode)
	//fmt.Println("app listen port:",config.AppConf.Port)

	//         192.168.233.130:8091
	//rs := polariser.PolarisProvider{}

	//port := 8322 //192.168.233.130:8500 10.107.9.140
	//register_client := consul.NewRegistryClient("192.168.233.130", 8500)
	//serviceId := fmt.Sprintf("%s", uuid.NewV4())
	//err := register_client.Register("10.107.9.140", 8088, "ginCloud", []string{"index"}, serviceId)
	//if err != nil {
	//	zap.S().Panic("服务注册失败:", err.Error())
	//}
	//zap.S().Debugf("启动服务器, 端口： %d", 8088)

	// PolarisProvider is an example of provider

	//Register()

	initArgs()
	flag.Parse()
	if len(namespace) == 0 || len(service) == 0 {
		log.Print("namespace and service are required")
		return
	}
	provider, err := polaris.NewProviderAPI()
	// 或者使用以下方法,则不需要创建配置文件
	// provider, err = api.NewProviderAPIByAddress("127.0.0.1:8091")

	if err != nil {
		log.Fatalf("fail to create providerAPI, err is %v", err)
	}
	defer provider.Destroy()

	svr := &PolarisProvider{
		provider:  provider,
		namespace: namespace,
		service:   service,
	}

	svr.Run()
	_ = route.Run(config.AppConf.Port)
	svr.runMainLoop()

	// register_client := register.NewRegistryClient(*RegistryClient,error)
	//接收终止信号
	//quit := make(chan os.Signal)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//<-quit
	//if err = register_client.DeRegister(serviceId); err != nil {
	//	zap.S().Info("注销失败:", err.Error())
	//} else {
	//	zap.S().Info("注销成功:")
	//}

}
