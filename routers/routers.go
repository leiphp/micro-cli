package routers

import (
	"github.com/gin-gonic/gin"
	"micro-cli/web/controllers"
	"net/http"
)

var (
	userController 	*controllers.UserController
	goodsController *controllers.GoodsController
)

//初始化所以控制器结构体
func initControllerStruct() {
	//用户控制器
	userController = controllers.NewUserController()
	//商品控制器 grpc调用
	goodsController = controllers.NewGoodsController()
}


// 初始化路由,service有多个还需要再次封装
func NewGinRouter() *gin.Engine {
	ginRouter := gin.Default()
	//ginRouter.Use(InitMiddleware(testService),ErrorMiddleware())
	//初始化控制器结构体
	initControllerStruct()
	//grpc路由
	v1Group := ginRouter.Group("/v1")
	{
		v1Group.Handle("POST","/goods", goodsController.Detail)
	}

	//普通路由
	v2 :=ginRouter.Group("/v2")
	{
		v2.GET("/user/:id",  userController.Detail)
		//v2.GET("/goods/:id",  goodsController.Detail)
	}

	v3 := ginRouter.Group("/check")
	{
		v3.GET("/health", func(c *gin.Context) {
			//c.String(http.StatusOK, "hello word")
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"data": []int{},
				"msg": "service is ok",
			})
		})
	}
	return ginRouter
}