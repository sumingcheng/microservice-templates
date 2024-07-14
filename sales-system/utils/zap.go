package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"microservice/sales-system/config"
)

// NewLogger 根据 LoggerConfig 创建一个新的 zap.Logger 实例
func NewLogger(c *config.LogConfig) (*zap.Logger, error) {
	env := config.GetEnv(config.FileConfig["ENV"])

	lumberLogger := &lumberjack.Logger{
		Filename:   c.FilePath,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
		Compress:   c.Compress,
	}

	defer func() {
		if err := lumberLogger.Close(); err != nil {
			log.Printf("Failed to close lumberjack logger: %v", err)
		}
	}()

	var cfg zapcore.EncoderConfig

	if env {
		cfg = zap.NewDevelopmentEncoderConfig()
	} else {
		cfg = zap.NewProductionEncoderConfig()
	}

	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewConsoleEncoder(cfg)
	core := zapcore.NewCore(fileEncoder, zapcore.AddSync(lumberLogger), config.ZapLevel)
	logger := zap.New(core)

	defer func() {
		if err := logger.Sync(); err != nil {
			log.Printf("Failed to sync logger: %v", err)
		}
	}()

	zap.ReplaceGlobals(logger)

	return logger, nil
}
