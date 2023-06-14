package boot

import "go-egg/pkg/config"

func SetupConfig(env string) {
	config.LoadConfig(env)
}
