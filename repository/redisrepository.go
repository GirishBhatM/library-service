package repository

import (
	"fmt"

	"github.com/go-redis/redis"
)

const REDIS_ADDR = "localhost:6379"
const REDIS_PWD = ""
const REDIS_DB = 0

func GetClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDR,
		Password: REDIS_PWD,
		DB:       REDIS_DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
		panic("unable to build redis client")
	}
	return client
}
