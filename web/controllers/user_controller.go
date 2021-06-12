package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"micro-cli/Services"
	"micro-cli/libs"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService Services.UserInterfaceService
}

func NewUserController() *UserController {
	obj := &UserController{
		UserService: Services.NewUserService(),
	}
	return obj
}


//获取用户详情页
func (this *UserController) Detail(c *gin.Context){
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	result, err := this.UserService.GetUserInfo(int64(id))
	log.Printf("[用户控制器-userInfo返回数据]-[%s]", libs.StructToJson(result))
	if err != nil {
		c.JSON(http.StatusNotFound, libs.ReturnJson(404, "", nil))
		return
	}
	c.JSON(http.StatusOK, libs.ReturnJson(200, "", result))

}
