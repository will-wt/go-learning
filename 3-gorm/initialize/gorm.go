package initialize

import (
	"fmt"
	"go-learning/3-gorm/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitGorm() *gorm.DB {
	ds := global.AppConfig.Datasource

	mysqlConfig := mysql.Config{
		DSN: ds.Dsn(),
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("数据库连接失败：%s \n", err))
	}

	// 连接池配置
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(ds.MaxIdle)
	sqlDB.SetMaxOpenConns(ds.MaxPoolSize)
	sqlDB.SetConnMaxIdleTime(time.Duration(ds.IdleTimeout) * time.Millisecond)
	sqlDB.SetConnMaxLifetime(time.Duration(ds.MaxLifetime) * time.Millisecond)

	return db
}
