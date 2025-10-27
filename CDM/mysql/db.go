package mysql

import (
	"bluebell/global"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var DB *sqlx.DB

// Init 初始化数据库连接
func Init() {
	cfg := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	var err error
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.S().Fatalf("数据库连接失败: %v\n", err)
	}

	// 设置连接池参数
	DB.SetMaxOpenConns(20) // 最大连接数
	DB.SetMaxIdleConns(10) // 空闲连接数
	DB.SetConnMaxLifetime(30 * time.Minute)

	// 测试连接
	if err = DB.Ping(); err != nil {
		zap.S().Fatalf("数据库无法连接: %v\n", err)
	}

	zap.S().Info("✅ 数据库连接成功")
}

// Close 关闭连接
func Close() {
	if DB != nil {
		DB.Close()
		zap.S().Info("✅ 数据库连接已关闭")
	}
}
