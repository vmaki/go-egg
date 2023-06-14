package handler

import (
	"github.com/gin-gonic/gin"
	"go-egg/config"
	"go-egg/pkg/logger"
	"net/http"
)

type UserHandle struct {
	BaseHandle
}

func (h *UserHandle) Index(ctx *gin.Context) {
	logger.DebugString("UserHandle", "Index", "访问成功")

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Hello " + config.GlobalConfig.Name,
	})
}
