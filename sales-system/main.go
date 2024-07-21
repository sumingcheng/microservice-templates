package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"log"
	"microservice/sales-system/config"
	"microservice/sales-system/middleware"
	"microservice/sales-system/router"
	"microservice/sales-system/utils"
)

func main() {
	// Config
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

	// MySQL
	db, err := utils.DBConnect(cfg.DBConfig)
	if err != nil {
		zap.S().Fatalf("Failed to connect to MySQL: %v", err.Error())
		return
	}

	// 注册自定义验证规则
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		utils.RegisterCustomValidations(v)
	} else {
		zap.S().Fatalf("Failed to assert validator engine type")
		return
	}

	// GIN
	r := gin.Default()
	// Middleware
	r.Use(middleware.Cors(cfg.AllowOrigin))
	// Router
	apiRouter := r.Group("/v1")
	router.Category(db, apiRouter, &utils.CustomError{})
	router.Product(db, apiRouter, &utils.CustomError{})
	router.Sale(db, apiRouter, &utils.CustomError{})

	// Run
	err = r.Run(cfg.GinConfig.IP + ":" + cfg.GinConfig.Port)
	if err != nil {
		zap.S().Fatalf("Failed to start gin server: %v", err.Error())
		return
	}

	fmt.Println("Server is running on " + cfg.GinConfig.IP + ":" + cfg.GinConfig.Port)
	zap.S().Infof("Server is running on %s:%s", cfg.GinConfig.IP, cfg.GinConfig.Port)
}
