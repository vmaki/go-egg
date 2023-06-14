package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB    *gorm.DB
	SqlDB *sql.DB
)

// Connect todo 后期可以支持读写分离&多库
func Connect(dbConfig gorm.Dialector, _logger logger.Interface) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})

	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 sqlDB
	SqlDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}
