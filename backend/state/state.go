package state

import (
	"fmt"
	"github.com/svemat01/shelley/pkg"
	"github.com/svemat01/shelley/redisDB"
	"time"
)

func Setup(data pkg.MainData) func() {

	// Calling NewTicker method
	Ticker := time.NewTicker(2 * time.Second)

	// Creating channel using make
	// keyword
	mychannel := make(chan bool)

	go tickerLoop(Ticker, mychannel, data)

	return func() {
		// Calling Stop() method
		Ticker.Stop()

		// Setting the value of channel
		mychannel <- true

		// Printed when the ticker is turned off
		fmt.Println("Ticker is turned off!")
	}
}

func fetchState(data pkg.MainData) {
	redisDB.GetDevice(data.Redis, data.RedisContext, "12345")
}

func tickerLoop(ticker *time.Ticker, channel chan bool, data pkg.MainData) {
	// Using for loop
	for {

		// Select statement
		select {

		// Case statement
		case <-channel:
			return

		// Case to print current time
		case <-ticker.C:
			fetchState(data)
		}
	}
}
