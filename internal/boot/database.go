package boot

import (
	"errors"
	"fmt"
	"go-egg/config"
	"go-egg/pkg/database"
	"go-egg/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func SetupDB() {
	cf := config.GlobalConfig.DataBase

	var dialector gorm.Dialector

	switch cf.Driver {
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			cf.Username,
			cf.Password,
			cf.Host,
			cf.Port,
			cf.Database,
			cf.Charset,
		)
		dialector = mysql.New(mysql.Config{
			DSN: dsn,
		})
	default:
		panic(errors.New("暂时不支持该数据库连接"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dialector, logger.NewGormLogger())

	database.SqlDB.SetMaxOpenConns(cf.MaxOpenConnections)
	database.SqlDB.SetMaxIdleConns(cf.MaxIdleConnections)
	database.SqlDB.SetConnMaxLifetime(time.Duration(cf.MaxLifeSeconds) * time.Second)

	// database.DB.AutoMigrate(&model.User{}) // 自动迁移
}
