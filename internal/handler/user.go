package handler

import (
	"github.com/gin-gonic/gin"
	"go-egg/config"
	"net/http"
)

type UserHandle struct {
	BaseHandle
}

func (h *UserHandle) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Hello " + config.GlobalConfig.Name,
	})
}
