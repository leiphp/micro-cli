package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"micro-cli/Services"
	"time"
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

// 定义中间
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}