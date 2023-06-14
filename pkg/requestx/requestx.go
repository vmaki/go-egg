package requestx

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-egg/pkg/logger"
)

func Validate(ctx *gin.Context, req any) error {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			logger.WarnString("requestx", "Validate", err.Error())
			return errors.New("格式错误或者非法参数")
		}

		return errors.New("参数错误")
	}

	return nil
}
