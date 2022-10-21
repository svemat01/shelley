package redisDB

import (
	"github.com/go-redis/redis/v9"
	"log"
	"strings"
)

type Device struct {
	Name string `json:"name" validate:"required"`
	Ip   string `json:"ip" validate:"required,ip"`
	Type string `json:"type" validate:"required,shellyType"`
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

func GetDevice(key string) (Device, error) {
	val, err := RedisClient.HGetAll(RedisContext, "device:"+key).Result()

	if err != nil {
		log.Println(err)
		return Device{}, err
	}

	// If val is empty, the device doesn't exist
	if len(val) == 0 {
		return Device{}, redis.Nil
	}

	device := Device{
		Name: val["name"],
		Ip:   val["ip"],
		Type: val["type"],
	}

	return device, nil
}

func ExistsDevice(key string) (bool, error) {
	val, err := RedisClient.Exists(RedisContext, "device:"+key).Result()

	if err != nil {
		return false, err
	}

	return val == 1, nil
}

func CreateDevice(key string, device Device) error {
	_, err := RedisClient.HSet(
		RedisContext,
		"device:"+key,
		"name", device.Name,
		"ip", device.Ip,
		"type", device.Type,
	).Result()

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
