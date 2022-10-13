package pkg

import (
	"context"
	"github.com/go-redis/redis/v9"
	"github.com/sony/sonyflake"
)

type MainData struct {
	Redis        *redis.Client
	RedisContext context.Context
}

var SonyFlake *sonyflake.Sonyflake
