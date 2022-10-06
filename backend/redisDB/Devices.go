package redisDB

import (
	"context"
	"github.com/go-redis/redis/v9"
	"strings"
)

func GetDevices(Redis *redis.Client, RedisContext context.Context) ([]string, error) {
	val, err := Redis.Keys(RedisContext, "device:*").Result()
	if err != nil {
		return []string{}, err
	}

	var keys []string
	for _, s := range val {
		keys = append(keys, strings.Replace(s, "device:", "", 1))
	}

	return keys, nil
}

func GetDevice(Redis *redis.Client, RedisContext context.Context, key string) (map[string]string, error) {
	val, err := Redis.HGetAll(RedisContext, "device:"+key).Result()

	if err != nil {
		return map[string]string{}, nil
	}

	return val, nil
}
