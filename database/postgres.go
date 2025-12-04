package database

import (
	"fmt"
	"gin-postgre-project/config"
	"gin-postgre-project/models"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.AppConfig.DBHost, config.AppConfig.DBUser, config.AppConfig.DBPassword, config.AppConfig.DBName, config.AppConfig.DBPort, config.AppConfig.DBSSLMode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("PostgreSQL 连接成功! ")
	// 自动迁移
	DB.AutoMigrate(
		&models.User{},
		&models.IDCInfo{},
		&models.MachineInfo{},
		&models.BusinessInfo{},
		&models.NetworkInfo{},
		&models.VersionInfo{},
	)

	// 初始化用户（只在第一次运行时候插入）
	seedUsers()
}

func seedUsers() {
	adminPass := "abcd001002"
	bdkejPass := "aabbccdd0102"

	users := []models.User{
		{Username: "admin", PasswordHash: hashPassword(adminPass), Role: "admin"},
		{Username: "bdkj", PasswordHash: hashPassword(bdkejPass), Role: "user"},
	}

	for _, user := range users {
		var exists int64
		DB.Model(&models.User{}).Where("username = ?", user.Username).Count(&exists)
		if exists == 0 {
			DB.Create(&user)
			fmt.Printf("创建用户成功: %s\n", user.Username)
		} else {
			// 已存在就更新密码(方便反复跑)
			DB.Model(&models.User{}).Where("username = ?", user.Username).Update("password_hash", user.PasswordHash)
			fmt.Printf("更新用户密码成功: %s\n", user.Username)
		}
	}
}

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
