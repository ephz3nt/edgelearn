package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	fmt.Println("Go Redis Tutorial")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("name", "Elliot", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	val, err := client.Get("name").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
