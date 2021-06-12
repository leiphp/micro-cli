package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"micro-cli/Services"
)

//context传递rpc service
func InitMiddleware(testService Services.TestService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["testservice"] = testService//赋值
		context.Next()
	}
}

//统一异常处理
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func () {
			if r := recover();r!=nil {
				context.JSON(500,gin.H{"status":fmt.Sprintf("%s",r)})
				context.Abort()
			}
		}()
		context.Next()
	}
}