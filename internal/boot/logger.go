package boot

import (
	"go-egg/config"
	"go-egg/pkg/helper"
	"go-egg/pkg/logger"
)

func SetLogger() {
	cf := config.GlobalConfig.Log
	isLocal := helper.IsLocal()

	logger.InitLogger(cf.Level, cf.Type, cf.Filename, cf.MaxSize, cf.MaxAge, cf.MaxBackup, cf.Compress, isLocal)
}
