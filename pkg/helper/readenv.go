package helper

import (
	"os"
	"time"

	"github.com/spf13/cast"
)

func EnvString(key string) string {
	return os.Getenv(key)
}

func EnvInt(key string) int {
	return cast.ToInt(os.Getenv(key))
}

func EnvInt64(key string) int64 {
	return cast.ToInt64(os.Getenv(key))
}

func EnvBool(key string) bool {
	return cast.ToBool(os.Getenv(key))
}

func EnvTime(key string) time.Duration {
	return cast.ToDuration(os.Getenv(key))
}
