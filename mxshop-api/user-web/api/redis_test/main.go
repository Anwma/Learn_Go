package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "192.168.119.128", 6379),
	})
	value, err := rdb.Get(context.Background(), "13549871354").Result()
	if err == redis.Nil {
		fmt.Println("key 不存在")
	}
	fmt.Println(value)

}
