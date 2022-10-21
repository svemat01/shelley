package redisDB

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
	"os"
)

var RedisContext = context.Background()
var RedisClient *redis.Client

func Setup() {
	// Setup redis
	redisAddress := os.Getenv("REDIS_HOST")
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	log.Println("Connecting to redis")
	if _, err := RedisClient.Ping(RedisContext).Result(); err != nil {

		log.Fatal("Can't connect to redis")
	}
	log.Println("Connected to redis")
}
