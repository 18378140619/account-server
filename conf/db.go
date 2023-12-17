package conf

import (
	"account-server/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitDB() (*gorm.DB, error) {
	logMode := logger.Info
	if !viper.GetBool("mode.develop") {
		logMode = logger.Error
	}

	db, err := gorm.Open(mysql.Open(viper.GetString("db.dns")), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		return nil, err
	}

	//配置空闲和使用
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConn"))
	sqlDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConn"))
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&model.User{})

	return db, err
}
