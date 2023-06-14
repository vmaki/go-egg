package handler

import (
	"github.com/gin-gonic/gin"
	"go-egg/config"
	"go-egg/internal/dto"
	"go-egg/pkg/logger"
	"go-egg/pkg/requestx"
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

func (h *UserHandle) Register(ctx *gin.Context) {
	var req dto.UserRegisterReq
	if err := requestx.Validate(ctx, &req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册",
	})
}
