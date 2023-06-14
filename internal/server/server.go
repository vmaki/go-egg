package server

import (
	"github.com/gin-gonic/gin"
	"go-egg/internal/middleware"
	"net/http"
	"strings"
)

// RegisterGlobalMiddleWare 注册全局中间件
func RegisterGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middleware.Logger(),
		middleware.Recovery(),
	)
}

// Setup404Handler 处理 404 请求
func Setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		accept := c.GetHeader("Accept")

		if strings.Contains(accept, "text/html") {
			c.String(http.StatusNotFound, "该页面不存在")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"code": 404,
				"msg":  "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
