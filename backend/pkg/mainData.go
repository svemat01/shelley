package pkg

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type MainData struct {
	Redis        *redis.Client
	RedisContext context.Context
}
