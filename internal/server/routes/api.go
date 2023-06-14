package routes

import (
	"github.com/gin-gonic/gin"
	"go-egg/internal/handler"
)

// RegisterAPIRoutes 注册API相关路由
func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		userGroup := v1.Group("/users")
		{
			handle := new(handler.UserHandle)

			userGroup.GET("/", handle.Index)
			userGroup.POST("/register", handle.Register)
		}
	}

	r.GET("/500", func(ctx *gin.Context) {
		panic("致命错误")
	})
}
