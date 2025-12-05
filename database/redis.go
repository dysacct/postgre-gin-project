package database

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()                         // 创建一个上下文, 用于与 Redis 交互
	if err := RedisClient.Ping(ctx).Err(); err != nil { // 如果连接失败, 则返回错误
		log.Fatal("Redis 连接失败", err)
	}
	fmt.Println("Redis 连接成功") // 如果连接成功, 则打印连接成功
}
