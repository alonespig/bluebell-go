package global

import (
	"bluebell/model"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	cfg := Config.Mysql

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Fatalln("failed to connect database", err)
		return
	}
	DB = db
	DB.AutoMigrate(&model.User{},
		&model.Community{},
		&model.Post{})
}

func Close() {
	defer func() {
		sqlDB, _ := DB.DB()
		sqlDB.Close()
	}()
	defer func() {
		Redis.Close()
	}()
	defer func() {
		Producer.Shutdown()
	}()
}
