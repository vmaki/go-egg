package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"go-egg/config"
	"go-egg/internal/boot"
	diyValidator "go-egg/internal/dto/validator"
)

func init() {
	boot.SetupConfig("")
	boot.SetLogger()
	boot.SetupDB()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("is-phone-exist", diyValidator.IsPhoneExist)
	}
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
