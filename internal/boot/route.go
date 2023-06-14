package boot

import (
	"go-egg/internal/middleware"
	"go-egg/internal/server"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	registerGlobalMiddleWare(router)

	server.RegisterAPIRoutes(router)

	setup404Handler(router)
}

// registerGlobalMiddleWare 注册全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middleware.Logger(),
		middleware.Recovery(),
	)
}

// setup404Handler 处理 404 请求
func setup404Handler(router *gin.Engine) {
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
