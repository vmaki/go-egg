package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-egg/config"
	"time"
)

var v *viper.Viper

func init() {
	v = viper.New()

	v.SetConfigType("yml")
	v.AddConfigPath("./config")
	v.SetEnvPrefix("go-es")
	v.AutomaticEnv()
}

func LoadConfig(env string) {
	configName := "settings.yml"
	if env != "" {
		configName = fmt.Sprintf("settings.%s.yml", env)
	}

	// 读取配置文件
	v.SetConfigName(configName)
	if err := v.ReadInConfig(); err != nil {
		panic("启动失败，err: 读取配置文件 " + configName + " 失败. " + err.Error())
	}

	// 加载配置
	if err := v.Unmarshal(config.GlobalConfig); err != nil {
		panic("启动失败，err: 加载配置失败，" + err.Error())
	}

	// 监控配置文件，变更时重新加载，无需重启
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件 - 重新加载", time.Now().Format("2006-01-02 15:04:05"))

		// todo 如果加载失败，需要记录到日志
		if err := v.Unmarshal(config.GlobalConfig); err != nil {
			fmt.Println("配置文件 - 重新加载失败" + err.Error())
		}
	})
}
