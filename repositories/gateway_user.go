package repositories

import (
	"log"
	"micro-cli/datamodels"
	"micro-cli/initialize"
)

/*
	操作gateway_user表的接口定义

	作者名称：leixiaotian 创建时间：20210513
*/

type GatewayUserInterface interface {
	SelectInfo(id int64) (datamodels.GatewayUser, error) //获得用户信息0
}

//返回结构体对象
func NewGatewayUser() GatewayUserInterface {
	return &gatewayUser{}
}

//gatewayUser构体
type gatewayUser struct {
}

//获得用户信息
func (this *gatewayUser) SelectInfo(id int64) (datamodels.GatewayUser, error) {

	var userInfo datamodels.GatewayUser
	//redis礼物key
	//jmfMemberKey := ReturnRedisKey(API_CACHE_JMF_MEMBER, userId)
	//result, err := initialize.RedisCluster.Get(jmfMemberKey).Bytes()

	//读取数据库
	if err := initialize.MsqlDb.Where("id = ? ", id).Find(&userInfo).Error; err != nil {
		log.Printf("[获取网关用户信息失败]-[%s]", err.Error())
		return userInfo, err
	}

	return userInfo, nil
}
