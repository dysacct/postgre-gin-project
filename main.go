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

	_ "gin-postgre-project/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// 1. 加载配置
	config.LoadConfig()
	// 2. 连接数据库 + 自动迁移 + 初始化用户
	database.ConnectDB()
	// 3. 连接 Redis
	database.ConnectRedis()
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 公开路由
	r.POST("/api/login", handlers.Login)

	// 需要登录的路由组
	auth := r.Group("/api")
	auth.Use(middleware.AuthRequired()) // 使用中间件, 需要登录
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
		// 机器相关路由
		auth.GET("/machines", handlers.ListMachines)            // 获取机器列表
		auth.GET("/machine/:zbx_id", handlers.GetMachine)       // 获取单个机器
		auth.POST("/machine", handlers.CreateMachine)           // 创建机器
		auth.PUT("/machine/:zbx_id", handlers.UpdateMachine)    // 更新机器
		auth.DELETE("/machine/:zbx_id", handlers.DeleteMachine) // 删除机器
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
}
