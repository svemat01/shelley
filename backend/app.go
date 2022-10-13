package main

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
	"github.com/sony/sonyflake"
	"github.com/svemat01/shelley/api"
	"github.com/svemat01/shelley/pkg"
	"github.com/svemat01/shelley/server"
	"github.com/svemat01/shelley/state"
	"log"
	"os"
)

var ctx = context.Background()

func main() {
	// Setup Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup ID Generator
	pkg.SonyFlake = sonyflake.NewSonyflake(sonyflake.Settings{})

	// Setup redis
	redisAddress := os.Getenv("REDIS_HOST")
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal("Can't connect to redis")
	}

	data := &pkg.MainData{
		Redis:        rdb,
		RedisContext: ctx,
	}

	// Server initialization
	app := server.Create()

	stopStateLoop := state.Setup(data)
	defer stopStateLoop()

	// Api routes
	api.Setup(app, data)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}

	//time.Sleep(8 * time.Second)
}
