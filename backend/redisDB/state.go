package redisDB

import (
	"encoding/json"
	"github.com/svemat01/shelley/pkg"
	"log"
)

func SetDeviceState(deviceId string, state pkg.DeviceState) error {
	value, err := json.Marshal(state)

	if err != nil {
		return err
	}

	err = RedisClient.Set(RedisContext, "state:"+deviceId, value, 0).Err()

	log.Printf("Device: %s state: %s\n", deviceId, string(value))

	if err != nil {
		return err
	}

	return nil
}

func GetDeviceState(deviceId string) (pkg.DeviceState, error) {
	val, err := RedisClient.Get(RedisContext, "state:"+deviceId).Result()

	if err != nil {
		return pkg.DeviceState{}, err
	}

	var state pkg.DeviceState
	err = json.Unmarshal([]byte(val), &state)

	if err != nil {
		return pkg.DeviceState{}, err
	}

	return state, nil
}
