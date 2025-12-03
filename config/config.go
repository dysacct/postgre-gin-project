package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	JWTSecret  string
	JWTExpire  string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("没有发现 .env 文件")
	}

	AppConfig = Config{
		DBHost:     os.Getenv("DB_HOST", "localhost"),
		DBPort:     os.Getenv("DB_PORT", "5432"),
		DBUser:     os.Getenv("DB_USER", "postgres"),
		DBPassword: os.Getenv("DB_PASSWORD", "abcd001002"),
		DBName:     os.Getenv("DB_NAME", "machine_info"),
		DBSSLMode:  os.Getenv("DB_SSL_MODE", "disable"),
		JWTSecret:  os.Getenv("JWT_SECRET", "fallback-secret-change-me"),
		JWTExpire:  os.Getenv("JWT_EXPIRE", "72"),
	}
}
