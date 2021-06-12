package Services

/*
	提供关于用户会员服务

	作者名称：leixiaotian 创建时间：20210530
*/
type UserInterfaceService interface {
}

//初始化对象函数
func NewUserService() UserInterfaceService {
	return &userService{

	}
}

type userService struct {

}
