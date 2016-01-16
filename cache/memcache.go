package cache

import (
	"github.com/kklis/gomemcache"
	"github.com/thejackrabbit/aero/conf"
	"github.com/thejackrabbit/aero/panik"
	"time"
)

type memCache struct {
	SimpleKeyFormat
	mc   *gomemcache.Memcache
	host string
	port int
}

func NewMemcache(host string, port int) Cacher {
	// connection check
	serv, err := gomemcache.Connect(host, port)
	panik.On(err)

	return memCache{
		mc:   serv,
		host: host,
		port: port,
	}
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
	key = c.Format(key)
	c.mc.Set(key, data, 0, int64(expireIn.Seconds()))
}

func (c memCache) Get(key string) ([]byte, error) {

	var data []byte
	var err error

	key = c.Format(key)
	data, _, err = c.mc.Get(key)

	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func (c memCache) Delete(key string) error {
	key = c.Format(key)
	return c.mc.Delete(key)
}

func (c memCache) Close() {
	panik.On(c.mc.Close())
}
