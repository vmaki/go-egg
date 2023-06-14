package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go-egg/config"
	"go-egg/internal/boot"
)

func init() {
	boot.SetupConfig("")
	boot.SetLogger()
	boot.SetupDB()
}

// main todo 后期需要迁移到 cmd 目录下面
func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	boot.SetupRoute(router)

	// 运行服务
	err := router.Run(":" + cast.ToString(config.GlobalConfig.Port))
	if err != nil {
		panic(err.Error())
	}
}
