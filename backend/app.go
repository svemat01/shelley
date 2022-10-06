package main

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
	"github.com/svemat01/shelley/api"
	"github.com/svemat01/shelley/pkg"
	"github.com/svemat01/shelley/server"
	"github.com/svemat01/shelley/state"
	"log"
)

var ctx = context.Background()

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "jab-school:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	data := pkg.MainData{
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
