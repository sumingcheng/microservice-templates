package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"microservice/sales-system/config"
	"microservice/sales-system/middleware"
	"microservice/sales-system/utils"
)

func main() {
	cfg, err := config.Initialization()
	if err != nil {
		log.Fatal("Failed to initialize config: ", err.Error())
		return
	}
	// Logger
	_, err = utils.NewLogger(cfg.LogConfig)
	if err != nil {
		zap.S().Fatalf("Failed to initialize logger: %v", err.Error())
		return
	}
	// GIN
	r := gin.Default()
	// MySQL
	db, err := utils.DBConnect(cfg.DBConfig)
	if err != nil {
		zap.S().Fatalf("Failed to connect to MySQL: %v", err.Error())
		return
	}
	// Middleware
	r.Use(middleware.Cors(cfg.AllowOrigin))
	r.Use(middleware.Router(db, r))

	// Run
	err = r.Run(cfg.GinConfig.IP + ":" + cfg.GinConfig.Port)
	if err != nil {
		zap.S().Fatalf("Failed to start gin server: %v", err.Error())
		return
	}

	fmt.Println("Server is running on " + cfg.GinConfig.IP + ":" + cfg.GinConfig.Port)
	zap.S().Infof("Server is running on %s:%s", cfg.GinConfig.IP, cfg.GinConfig.Port)
}
