package main

import (
	"go.uber.org/zap"
	"microservice/sales-system/config"
	"microservice/sales-system/model"
	"microservice/sales-system/utils"
)

func main() {
	cfg, err := config.Initialization()
	if err != nil {
		zap.S().Fatalf("Failed to initialize config: %v", err.Error())
		return
	}

	// Logger
	_, err = utils.NewLogger(cfg.LogConfig)

	db, err := utils.DBConnect(cfg.DBConfig)

	if err != nil {
		zap.S().Fatalf("Failed to connect to MySQL: %v", err.Error())
		return
	}

	defer func() {
		err := utils.DBClose(db)
		if err != nil {
			zap.S().Errorf("Failed to close MySQL connection: %v", err)
		}
	}()

	modules := []any{
		&model.Product{},
		&model.Category{},
		&model.Sale{},
	}

	err = db.AutoMigrate(modules...)

	if err != nil {
		zap.S().Fatalf("Failed to migrate tables: %v", err.Error())
		return
	}
}
