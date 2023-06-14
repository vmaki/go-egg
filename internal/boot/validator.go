package boot

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	myValidator "go-egg/internal/dto/validator"
)

func InitValidator() {
	var rules = map[string]validator.Func{
		"is-phone-exist": myValidator.IsPhoneExist,
	}

	// 加载规则
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for tag, v2 := range rules {
			_ = v.RegisterValidation(tag, v2)
		}
	}
}
