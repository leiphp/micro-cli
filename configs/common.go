/*
@Time : 2021/04/12 14:45
@Author : LeiXiaoTian
@File : common.go
@Software: GoLand
*/
package configs

//定义公用常量
const (

	//一些常量
	Uniacid                       = 3    //挂接的公众号$uniacid
	ExchangeGoodsId               = 226  //兑换使用的商品ID
	CheckinGoodsId                = 227  //签到使用的商品ID
	RechargepackageAndroidGoodsId = 223  //安卓平台充值套餐使用的商品ID
	RechargepackageIosGoodsId     = 224  //ios平台充值套餐使用的商品ID
	RechargepackageWebGoodsId     = 225  //web平台充值套餐使用的商品ID
	VipPackageCategoryId          = 1214 //会员套餐使用的分类ID
	RechargePackageCategoryId     = 1215 //充值套餐使用的分类ID
	GiftsCategoryId               = 1217 //礼物分类Id
	DefaultPage                   = 1    //默认第一页
	DefaultPerPage                = 10   //默认一页10条记录


	Page    = 1
	PerPage = 10

)

//定义消息容器
var MsgCode map[int]string = map[int]string{
	200: "成功",
	400: "客户端请求的语法错误，服务器无法理解",
	403: "服务器理解请求客户端的请求，但是拒绝执行此请求",
	404: "请求的资源不存在",
	405: "客户端请求中的方法被禁止",
	408: "服务器等待客户端发送的请求时间过长，超时",
	500: "内部服务器错误",
	501: "服务器不支持请求的功能，无法完成请求",

	//系统相关
	500100: "系统错误",
}
