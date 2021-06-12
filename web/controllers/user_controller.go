package controllers

import (
	"github.com/gin-gonic/gin"
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
	c.JSON(http.StatusOK, libs.ReturnJson(200, "", id))
}

func (this *UserController) Health(c *gin.Context){
	c.JSON(http.StatusOK, libs.ReturnJson(200, "", "ok"))
}