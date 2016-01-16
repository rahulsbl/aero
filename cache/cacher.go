package cache

import (
	"github.com/thejackrabbit/aero/conf"
	"github.com/thejackrabbit/aero/panik"
	"strings"
	"time"
)

type KeyFormatter interface {
	Format(key string) string
}

type SimpleKeyFormat struct {
}

func (k SimpleKeyFormat) Format(key string) string {
	return strings.Replace(key, " ", "-", -1)
}

type Cacher interface {
	KeyFormatter
	Set(key string, data []byte, expireIn time.Duration)
	Get(key string) ([]byte, error)
	Delete(key string) error
	Close()
}

func FromConfig(container string) (out Cacher) {

	cType := conf.String("", container, "type")
	panik.If(cType == "", "cache type is not specified")

	switch cType {
	case "memcache":
		out = MemcacheFromConfig(container)

	case "inmem", "inmemory":
		out = InMemoryFromConfig(container)

	case "debug":
		out = DebugFromConfig(container)

	case "redis":
		out = RedisFromConfig(container)

	default:
		panik.Do("Unknown cache provider: %s", cType)
	}

	return out
}
