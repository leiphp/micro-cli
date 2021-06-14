package controllers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"log"
	"micro-cli/Services"
	"micro-cli/initialize"
	"micro-cli/libs"
	"net/http"
	"strconv"
)

type GoodsController struct {

	GoodsService Services.GoodsInterfaceService
	//RpcGoodsService Services.GoodsServiceClient
	RpcTestService Services.TestService
}

func NewGoodsController() *GoodsController {
	obj := &GoodsController{
		GoodsService: Services.NewGoodsService(),
		//RpcGoodsService: Services.NewGoodsServiceClient(initialize.GrpcConn),
		RpcTestService: Services.NewTestService("go.micro.blog",initialize.Client),//指定grpc服务名
	}
	return obj
}


//获取商品详情页
func (this *GoodsController) Detail(c *gin.Context){
	type Goods struct {
		Id int64 `json:"id"`
		Name string `json:"name"`
	}
	goods := Goods{}
	c.BindJSON(&goods)

	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	initialize.RusLog.Infof("[用户控制器-http请求数据]-[%d]", id)
	log.Printf("[用户控制器-http请求数据query]-[%d]", id)
	log.Printf("[用户控制器-http请求数据json]-[%d]", goods.Id)
	//熔断代码改造
	//第一步：配置config
	configA := hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  0,
		RequestVolumeThreshold: 0,
		SleepWindow:            0,
		ErrorPercentThreshold:  0,
	}
	//第二步：配置command
	hystrix.ConfigureCommand("getprods",configA)
	//第三步：执行使用Do方法，同步
	var goodsRes *Services.TestResponse
	err := hystrix.Do("getprods",func() error {
		var err error
		goodsRes, err= this.RpcTestService.Call(context.Background(), &Services.TestRequest{Id: int32(goods.Id)})
		return err
	}, func(e error) error {//降级方法
		goodsRes,e = this.defaultProds(c)
			return e
		})
	//goodsRes, err := this.RpcTestService.Call(context.Background(), &Services.TestRequest{Id: int32(goods.Id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError,  libs.ReturnJson(500, err.Error(), gin.H{}))
		return
	}

	//c.JSON(http.StatusOK, libs.ReturnJson(200, "", result))
	c.JSON(http.StatusOK, libs.ReturnJson(200, "", gin.H{"goods_id": id, "remark": "库存数量等于goods_id乘以10", "goods": goodsRes}))
}

//降级方法
func (this *GoodsController) defaultProds(c *gin.Context) (*Services.TestResponse,error){
	res := &Services.TestResponse{}
	res.Data = "降级"+strconv.Itoa(int(500))
	return res, nil
}