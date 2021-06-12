package controllers

import (
	"context"
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
	log.Printf("[用户控制器-http请求数据query]-[%d]", id)
	log.Printf("[用户控制器-http请求数据json]-[%d]", goods.Id)

	goodsRes, err := this.RpcTestService.Call(context.Background(), &Services.TestRequest{Id: int32(goods.Id)})
	if err != nil {
		log.Printf("[用户控制器-http请求数据]-[%d]", id)
		c.JSON(http.StatusInternalServerError,  libs.ReturnJson(500, "", gin.H{}))
		return
	}

	//c.JSON(http.StatusOK, libs.ReturnJson(200, "", result))
	c.JSON(http.StatusOK, libs.ReturnJson(200, "", gin.H{"goods_id": id, "remark": "库存数量等于goods_id乘以10", "goods": goodsRes}))
}
