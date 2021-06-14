package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/metadata"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"micro-cli/configs"
	"micro-cli/initialize"
	"micro-cli/routers"
)

//装饰器拦截
type logWrapper struct {
	client.Client
}

//装饰器执行
func(this *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error{
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return this.Client.Call(ctx,req,rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}


func main(){
	//获得配置对象
	Yaml := configs.InitConfig()
	initialize.Init(Yaml)
	//micro-cli作为内部服务，注册服务到etcd
	//go-micro v3启动http服务，参考https://github.com/asim/go-micro/tree/master/plugins/server/http
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	//创建http server
	srv := httpServer.NewServer(
		server.Name("go.micro.http"),
		server.Address(":8002"),
	)

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte(`hello world`))
	//})

	hd := srv.NewHandler(routers.NewGinRouter())

	srv.Handle(hd)

	service := micro.NewService(
		micro.WrapClient(NewLogWrapper),
		micro.Server(srv),
		micro.Registry(etcdReg),
	)

	//调grpc服务
	//testService := Services.NewTestService("go.micro.blog",service.Client())//指定grpc服务名

	service.Init()
	service.Run()
}

