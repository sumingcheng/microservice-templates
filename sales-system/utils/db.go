package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"microservice/sales-system/config"
	"os"
	"time"
)

func DBConnect(c *config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.PassWord, c.Host, c.Port, c.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second, // 慢查询阈值
				LogLevel:      logger.Info, // 日志级别
				Colorful:      true,        // 启用彩色打印
			},
		),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Prefix, // 表名前缀
			SingularTable: true,     // 表名单数形式
		},
		PrepareStmt: true, // 启用PrepareStmt, SQL预编译，提高查询效率
	})

	return db, err
}

func DBClose(db *gorm.DB) (err error) {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}
	return nil
}
