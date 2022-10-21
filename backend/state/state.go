package state

import (
	"fmt"
	"github.com/svemat01/shelley/redisDB"
	"github.com/svemat01/shelley/shelly"
	"log"
	"sync"
	"time"
)

func Setup() func() {

	// Calling NewTicker method
	Ticker := time.NewTicker(5 * time.Second)

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

			wg := sync.WaitGroup{}
			wg.Add(len(devices))
			log.Printf("Found %d devices", len(devices))

			for _, deviceId := range devices {
				log.Printf("Fetching state for device %s", deviceId)
				go func(deviceId string) {
					device, err := redisDB.GetDevice(deviceId)

					if err != nil {
						fmt.Println(err)
						wg.Done()
						return
					}

					state, err := shelly.GetDeviceState(device)

					err = redisDB.SetDeviceState(deviceId, state)

					if err != nil {
						fmt.Println(err)
						wg.Done()
						return
					}

					wg.Done()
				}(deviceId)
			}

			log.Printf("Waiting for all devices to be fetched")
			wg.Wait()
		}
	}
}
