package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	PythonHost string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config") // tên file: app.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config") // thư mục chứa file config

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Lỗi đọc config: %v", err)
	}

	AppConfig = Config{
		ServerPort: viper.GetString("server.port"),
		DBHost:     viper.GetString("database.host"),
		DBUser:     viper.GetString("database.user"),
		DBPassword: viper.GetString("database.password"),
		DBName:     viper.GetString("database.name"),
		DBPort:     viper.GetString("database.port"),
		PythonHost: viper.GetString("python.host"),
	}
}
