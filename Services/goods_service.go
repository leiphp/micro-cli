package Services

import (
	"micro-cli/repositories"
)

/*
	提供关于商品服务

	作者名称：leixiaotian 创建时间：20210604
*/
type GoodsInterfaceService interface {
	GetInfo(id int64) (interface{},error) //获取商品详情
}

//初始化对象函数
func NewGoodsService() GoodsInterfaceService {
	return &goodsService{
		gatewayUserService:          repositories.NewGatewayUser(),
	}
}

type goodsService struct {
	gatewayUserService 			    repositories.GatewayUserInterface           //网关会员服务
}

//获取商品信息
func (this *goodsService) GetInfo(id int64) (interface{},error){

	//userInfo, err := this.gatewayUserService.SelectInfo(id)
	//initialize.IrisLog.Infof("[商品服务-userInfo数据]-[%s]", libs.StructToJson(userInfo))
	//if err != nil {
	//	initialize.IrisLog.Errorf("[商品服务-获取用户信息失败]-[%s]", err.Error())
	//	return 3006, err
	//}
	//userInfo.CreateDate = time.Unix(userInfo.CreateTime, 0).Format("2006-01-02 15:04:05")
	//return userInfo, nil
	return nil,nil
}