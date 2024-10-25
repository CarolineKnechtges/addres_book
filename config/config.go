// Package config/config.go
package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

var DB *gorm.DB

type Config struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	DBName    string `mapstructure:"name"`
	Charset   string `mapstructure:"charset"`
	ParseTime bool   `mapstructure:"parseTime"`
	Loc       string `mapstructure:"loc"`
}

func LoadConfig() (Config, error) {
	var config Config

	// 获取项目的根目录
	rootDir, err := os.Getwd()
	if err != nil {
		return config, fmt.Errorf("failed to get working directory: %v", err)
	}

	// 设置配置文件的路径（使用绝对路径）
	configPath := filepath.Join(rootDir, "config")
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".") // 也可添加相对路径查找

	viper.SetConfigName("config_dev")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("config_dev.yml not found, falling back to config.yml")

		// 尝试读取正式配置
		viper.SetConfigName("config")
		if err := viper.ReadInConfig(); err != nil {
			return config, fmt.Errorf("error reading config file, %s", err)
		}
	}

	if err := viper.UnmarshalKey("database", &config); err != nil {
		return config, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return config, nil
}

func ConnectDatabase() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 生成 DSN（数据源名称）
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.Charset,
		config.ParseTime,
		config.Loc,
	)

	// 连接数据库
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = database
}
