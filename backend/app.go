package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/svemat01/shelley/api"
	"github.com/svemat01/shelley/redisDB"
	"github.com/svemat01/shelley/server"
	"github.com/svemat01/shelley/state"
)

func main() {
	// Setup Env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
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
