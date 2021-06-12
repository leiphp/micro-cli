package Services

import (
	"log"
	"micro-cli/libs"
	"micro-cli/repositories"
	"time"
)

/*
	提供关于用户会员服务

	作者名称：leixiaotian 创建时间：20210530
*/
type UserInterfaceService interface {
	GetUserInfo(id int64) (interface{},error) //获取用户详情
}

//初始化对象函数
func NewUserService() UserInterfaceService {
	return &userService{
		gatewayUserService:          repositories.NewGatewayUser(),
	}
}

type userService struct {
	gatewayUserService 			    repositories.GatewayUserInterface           //网关会员服务
}

//获取用户信息
func (this *userService) GetUserInfo(id int64) (interface{},error){
	userInfo, err := this.gatewayUserService.SelectInfo(id)
	log.Printf("[用户服务-userInfo数据]-[%s]", libs.StructToJson(userInfo))
	if err != nil {
		log.Printf("[用户服务-获取用户信息失败]-[%s]", err.Error())
		return 3006, err
	}
	userInfo.CreateDate = time.Unix(userInfo.CreateTime, 0).Format("2006-01-02 15:04:05")
	return userInfo, nil
}