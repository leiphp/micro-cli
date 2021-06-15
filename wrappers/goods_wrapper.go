package wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/v3/client"
	"micro-cli/Services"
	"strconv"
)

type ProdsWrapper struct {
	client.Client
}

func(this *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error{
	cmdName := req.Service()+"."+req.Endpoint()
	configA := hystrix.CommandConfig{
		Timeout:                2000,
		MaxConcurrentRequests:  0,
		RequestVolumeThreshold: 2,//请求阀值
		SleepWindow:            5000,//熔断再次监测是否开启
		ErrorPercentThreshold:  50,//错误百分比
	}
	hystrix.ConfigureCommand(cmdName,configA)
	return hystrix.Do(cmdName, func() error {
		return this.Client.Call(ctx,req,rsp)
	}, func(e error) error {
		//defaultProds(rsp)
		defaultData(rsp)
		return nil
	})
}

func NewProdsWrapper (c client.Client) client.Client {
	return &ProdsWrapper{c}
}

//测试方法
//func newProd(id int32,pname string) *Services.ProdModel {
//	return &Services.ProdModel{ProdID: id,ProdName:pname}
//}

//通用降级方法
func defaultData(rsp interface{}) {
	//switch t:=rsp.(type) {
	switch rsp.(type) {
	case *Services.TestResponse:
		defaultProds(rsp)
	//case *Services.ProdDetailResponse:
	//	t.Data=newProd(10,"降级商品")

	}
}

//商品列表降级方法
func defaultProds(rsp interface{}) {
	result := rsp.(*Services.TestResponse)
	result.Data="降级"+strconv.Itoa(int(500))
	//return result, nil
}
