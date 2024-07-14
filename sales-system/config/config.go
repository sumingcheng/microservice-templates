package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
	"log"
)

type GinConfig struct {
	IP   string `mapstructure:"ip"`
	Port string `mapstructure:"port"`
}

type LogConfig struct {
	FilePath   string `mapstructure:"file_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type TransConfig struct {
	Locale string `mapstructure:"locale"`
}

type DBConfig struct {
	User     string `mapstructure:"user"`
	PassWord string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
	Prefix   string `mapstructure:"prefix"`
}

type AllowOrigin struct {
	Origins []string `mapstructure:"origins"`
}

type Config struct {
	*GinConfig   `mapstructure:"gin"`
	*DBConfig    `mapstructure:"mysql"`
	*TransConfig `mapstructure:"translator"`
	*LogConfig   `mapstructure:"log"`
	*AllowOrigin `mapstructure:"allow_origin"`
}

var cfg = new(Config)
var ZapLevel = zapcore.InfoLevel

func Initialization() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetConfigFile(getEnvFile())
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatal("解析配置文件失败", err)
			return
		}
	})

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func getEnvFile() string {
	var configFile string
	isDEV := viper.GetString("ENV") == "DEV"

	if isDEV {
		configFile = fmt.Sprintf("./sales-system/config/config-prod.yaml")
	} else {
		configFile = fmt.Sprintf("./sales-system/config/config-dev.yaml")
	}
	fmt.Println("configFile: ", configFile)
	return configFile
}
