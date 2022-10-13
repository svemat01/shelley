package pkg

import (
	"github.com/sony/sonyflake"
)

var SonyFlake *sonyflake.Sonyflake

func init() {
	SonyFlake = sonyflake.NewSonyflake(sonyflake.Settings{})
}
