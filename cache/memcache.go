package cache

import (
	"github.com/kklis/gomemcache"
	"github.com/thejackrabbit/aero/conf"
	"github.com/thejackrabbit/aero/panik"
	"time"
)

type memCache struct {
	mc *gomemcache.Memcache
}

func NewMemcache(host string, port int) Cacher {
	serv, err := gomemcache.Connect(host, port)
	panik.On(err)

	return memCache{
		mc: serv,
	}
	// TODO: close port on destruction
}

// memcache:
// - host
// - port
func MemcacheFromConfig(container string) Cacher {
	host := conf.String("", container, "host")
	panik.If(host == "", "memcache host not specified")

	port := conf.Int(0, container, "port")
	panik.If(port == 0, "memcache port not specified")

	return NewMemcache(host, port)
}

func (c memCache) Set(key string, data []byte, expireIn time.Duration) {
	c.mc.Set(prepareKey(key), data, 0, int64(expireIn.Seconds()))
}

func (c memCache) Get(key string) ([]byte, error) {
	data, _, err := c.mc.Get(prepareKey(key))
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
