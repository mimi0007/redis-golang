package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("Hello World of redis!")

	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
	)

	ctx := context.Background()

	fmt.Println(redisClient.Ping(ctx)) //this will print ping: PONG if the redis client is connected successfully

	err := redisClient.Set(ctx, "key", "VALUE", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisClient.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
