// main.go
package main

import (
	"fmt"
	"gin-postgre-project/config"
	"gin-postgre-project/database"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 1. 加载配置
	config.LoadConfig()

	// 2. 连接数据库 + 自动迁移 + 初始化用户
	database.ConnectDB()

	fmt.Println("阶段二全部完成！数据库连接正常，用户已就绪")
	fmt.Println("按 Ctrl+C 安全退出程序")

	// 优雅退出：捕获 Ctrl+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop // 阻塞等待信号

	fmt.Println("\n正在关闭数据库连接...")
	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Fatal("获取 sql.DB 失败:", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Println("关闭数据库连接失败:", err)
	} else {
		fmt.Println("数据库连接已关闭，程序安全退出")
	}
}
