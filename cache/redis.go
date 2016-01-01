package cache

import (
	"fmt"
	"github.com/thejackrabbit/aero/conf"
	"github.com/thejackrabbit/aero/panik"
	"gopkg.in/redis.v3"
	"time"
)

type redisStore struct {
	r *redis.Client
}

func NewRedis(host string, port int, db int) Cacher {
	serv := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
		DB:   int64(db),
	})
	_, err := serv.Ping().Result()
	panik.On(err)

	return redisStore{
		r: serv,
	}
	// TODO: close port on destruction
}

// redis:
// - host
// - port
// - db
func RedisFromConfig(container string) Cacher {
	host := conf.String("", container, "host")
	panik.If(host == "", "redis host not specified")

	port := conf.Int(0, container, "port")
	panik.If(port == 0, "redis port not specified")

	db := conf.Int(0, container, "db")

	return NewRedis(host, port, db)
}

func (rd redisStore) Set(key string, data []byte, expireIn time.Duration) {
	rd.r.Set(prepareKey(key), data, expireIn)
}

func (rd redisStore) Get(key string) ([]byte, error) {
	data, err := rd.r.Get(prepareKey(key)).Bytes()
	if err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func (rd redisStore) Close() {
	panik.On(rd.r.Close())
}
