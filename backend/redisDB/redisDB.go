package redisDB

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
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
	continueChecking := true
	tried := 0

	for continueChecking {
		_, err := RedisClient.Ping(RedisContext).Result()
		if err != nil {
			log.Println("Error connecting to redis")
			log.Println(err)
		} else {
			log.Println("Connected to redis")
			continueChecking = false
			break
		}

		if tried > 10 {
			log.Println("Tried to many times, Exiting...")
			os.Exit(1)
		}

		tried++

		log.Println("Trying again in 10 seconds")
		time.Sleep(10 * time.Second)
	}
}
