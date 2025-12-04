// main.go
package main

import (
	"context"
	"fmt"
	"gin-postgre-project/config"
	"gin-postgre-project/database"
	"gin-postgre-project/handlers"
	"gin-postgre-project/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载配置
	config.LoadConfig()
	// 2. 连接数据库 + 自动迁移 + 初始化用户
	database.ConnectDB()
	// 3. 连接 Redis
	database.ConnectRedis()
	r := gin.Default()

	// 公开路由
	r.POST("/api/login", handlers.Login)

	// 需要登录的路由组
	auth := r.Group("/api")
	auth.Use(middleware.AuthRequired())
	{
		auth.GET("/ping", func(c *gin.Context) {
			username, _ := c.Get("username")
			c.JSON(200, gin.H{
				"message":  "pong",
				"login_as": username,
			})
		})

		// 管理员专属路由
		admin := auth.Group("/admin")
		admin.Use(middleware.AdminRequired())
		{
			admin.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "管理员面板"})
			})
		}
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		fmt.Println("API 服务器启动 -> http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("服务器启动失败:", err)
		}
	}()

	// 优雅退出：捕获 Ctrl+C
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stop // 阻塞等待信号

	fmt.Println("\n正在关闭服务...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	fmt.Println("服务已关闭，程序安全退出")
	// fmt.Println("\n正在关闭数据库连接...")
	// sqlDB, err := database.DB.DB()
	// if err != nil {
	// 	log.Fatal("获取 sql.DB 失败:", err)
	// }
	// if err := sqlDB.Close(); err != nil {
	// 	log.Println("关闭数据库连接失败:", err)
	// } else {
	// 	fmt.Println("数据库连接已关闭，程序安全退出")
	// }
}
