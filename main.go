package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type student struct {
	Name  string
	Class int32
	Age   uint
}

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

	newStudent := student{
		Name:  "Alice",
		Class: 10,
		Age:   16,
	}

	json, err := json.Marshal(newStudent)
	if err != nil {
		panic(err)
	}

	err = redisClient.Set(ctx, "key", json, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := redisClient.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(val)
}
