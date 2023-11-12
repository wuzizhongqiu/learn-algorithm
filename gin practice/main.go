package main

import (
	"github.com/gin-gonic/gin"
)

func initRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

}

func main() {
	r := gin.Default()

	r.GET("/hello", hello)

	r.Run(":9090")
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello gin",
	})
}
