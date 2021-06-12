package main

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
	"micro-cli/Services"
	"micro-cli/web/controllers"
)


var (
	userController 	*controllers.UserController
	//goodsController *controllers.GoodsController
)

//初始化所以控制器结构体
func initControllerStruct() {
	//用户控制器
	userController = controllers.NewUserController()
	//商品控制器 grpc调用
	//goodsController = controllers.NewGoodsController()
}

func main(){
	//micro-cli作为内部服务，注册服务到etcd
	//go-micro v3启动http服务，参考https://github.com/asim/go-micro/tree/master/plugins/server/http
	etcdReg := etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))

	ginRouter := gin.Default()
	//创建http server
	srv := httpServer.NewServer(
		server.Name("go.micro.http"),
		server.Address(":8002"),
	)

	//mux := http.NewServeMux()
	//mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte(`hello world`))
	//})

	hd := srv.NewHandler(ginRouter)

	srv.Handle(hd)

	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(etcdReg),
	)

	//调grpc服务
	testService := Services.NewTestService("go.micro.blog",service.Client())//指定grpc服务名

	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST","/prods", func(ginCtx *gin.Context) {
			var testReq Services.TestRequest
			err := ginCtx.Bind(&testReq)
			if err != nil {
				ginCtx.JSON(500, gin.H{"status":err.Error()})

			}else{
				testRes, _ := testService.Call(context.Background(),&testReq)
				ginCtx.JSON(200,gin.H{"data":testRes.Data})
			}

		})
	}

	//普通路由
	v2 :=ginRouter.Group("/v2")
	{
		v2.POST("/health", userController.Health)
		v2.GET("/user/:id",  userController.Detail)
		//v2.GET("/goods/:id",  goodsController.Detail)
	}

	service.Init()
	service.Run()
}

