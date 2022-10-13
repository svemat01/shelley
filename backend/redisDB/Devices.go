package redisDB

import (
	"strings"
)

type Device struct {
	Name string `json:"name" validate:"required"`
	Ip   string `json:"ip" validate:"required,ip"`
}

func GetDevices() ([]string, error) {
	val, err := RedisClient.Keys(RedisContext, "device:*").Result()
	if err != nil {
		return []string{}, err
	}

	var keys []string
	for _, s := range val {
		keys = append(keys, strings.Replace(s, "device:", "", 1))
	}

	return keys, nil
}

func GetDevice(key string) (map[string]string, error) {
	val, err := RedisClient.HGetAll(RedisContext, "device:"+key).Result()

	if err != nil {
		return map[string]string{}, nil
	}

	return val, nil
}

func ExistsDevice(key string) (bool, error) {
	val, err := RedisClient.Exists(RedisContext, "device:"+key).Result()

	if err != nil {
		return false, err
	}

	return val == 1, nil
}

func CreateDevice(key string, device Device) error {
	_, err := RedisClient.HSet(RedisContext, "device:"+key, "name", device.Name, "ip", device.Ip).Result()

	if err != nil {
		return err
	}

	return nil
}

func DeleteDevice(key string) error {
	_, err := RedisClient.Del(RedisContext, "device:"+key).Result()

	if err != nil {
		return err
	}

	return nil
}
