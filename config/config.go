// Package config/config.go
package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
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

	// 设置默认配置文件
	viper.SetConfigName("config_dev") // 开发配置文件名 (不需要扩展名)
	viper.SetConfigType("yaml")       // 配置文件类型
	viper.AddConfigPath(".")          // 在项目根目录下查找配置文件

	// 尝试读取开发配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("config_dev.yml not found, falling back to config.yml")

		// 尝试读取正式配置文件
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
