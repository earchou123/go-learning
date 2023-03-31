package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func initRedis(address string) (client *redis.Client) {
	client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // 无密码连接
		DB:       0,  // 默认数据库
	})

	// 使用 ping 命令测试连接是否正常
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("连接 Redis 失败：", err)
		return
	}
	fmt.Println("连接 Redis 成功：", pong)

	return
}
