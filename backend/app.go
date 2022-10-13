package main

import (
	"github.com/joho/godotenv"
	"github.com/svemat01/shelley/api"
	"github.com/svemat01/shelley/redisDB"
	"github.com/svemat01/shelley/server"
	"github.com/svemat01/shelley/state"
	"log"
)

func main() {
	// Setup Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Setup Redis
	redisDB.Setup()

	// Server initialization
	app := server.Create()

	stopStateLoop := state.Setup()
	defer stopStateLoop()

	// Api routes
	api.Setup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}

	//time.Sleep(8 * time.Second)
}
