package state

import (
	"encoding/json"
	"fmt"
	"github.com/svemat01/shelley/redisDB"
	"github.com/svemat01/shelley/shelly"
	"log"
	"time"
)

type DeviceState struct {
	State      bool
	Brightness int
}

func Setup() func() {

	// Calling NewTicker method
	Ticker := time.NewTicker(2 * time.Second)

	// Creating channel using make
	// keyword
	mychannel := make(chan bool)

	go tickerLoop(Ticker, mychannel)

	return func() {
		// Calling Stop() method
		Ticker.Stop()

		// Setting the value of channel
		mychannel <- true

		// Printed when the ticker is turned off
		fmt.Println("Ticker is turned off!")
	}
}

func fetchState(device redisDB.Device) {

}

func tickerLoop(ticker *time.Ticker, channel chan bool) {
	// Using for loop
	for {

		// Select statement
		select {

		// Case statement
		case <-channel:
			return

		// Case to print current time
		case <-ticker.C:
			devices, err := redisDB.GetDevices()

			if err != nil {
				fmt.Println(err)
				return
			}

			for _, deviceId := range devices {
				device, err := redisDB.GetDevice(deviceId)

				if err != nil {
					fmt.Println(err)
					continue
				}

				response, err := shelly.GetDeviceState(device.Ip, device)

				jsonResponse, _ := json.Marshal(response)
				// TODO set state in redis
				log.Printf("Device: %s state: %s\n", deviceId, string(jsonResponse))
			}
		}
	}
}
