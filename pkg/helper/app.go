package helper

import "go-egg/config"

func IsLocal() bool {
	return config.GlobalConfig.Mode == "local"
}
