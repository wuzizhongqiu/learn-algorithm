package main

import (
	context "context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func initClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("connect redis fails")
		return
	}
}

func main() {
	initClient()
	fmt.Println("connect success")
	defer rdb.Close()
}
