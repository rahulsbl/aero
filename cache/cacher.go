package cache

import (
	"github.com/thejackrabbit/aero/conf"
	"github.com/thejackrabbit/aero/db/cstr"
	"github.com/thejackrabbit/aero/engine"
	"github.com/thejackrabbit/aero/key"
	"github.com/thejackrabbit/aero/panik"
	"strings"
	"time"
)

type Cacher interface {
	key.KeyFormatter
	Set(key string, data []byte, expireIn time.Duration)
	Get(key string) ([]byte, error)
	Delete(key string) error
	Close()
}

func NewCacher(container ...string) (out Cacher) {
	parent := strings.Join(container, ".")

	engn := conf.String("", parent, "engine")
	panik.If(engn == "", "cache engine is not specified")

	switch engn {
	case "memcache":
		{
			cnf := cstr.Memcache{}
			conf.Struct(&cnf, parent)
			out = engine.NewMemcache(cnf.Host, cnf.Port)
		}

	case "inproc", "inproccache":
		{
			expiry := conf.String("1h", parent, "lifetime")
			life, err := time.ParseDuration(expiry)
			panik.On(err)
			out = engine.NewInProcCache(life)
		}

	case "log":
		{
			dir := conf.String("", parent, "dir")
			inner := NewCacher(parent, "inner")
			out = NewCacheLogger(dir, inner)
		}

	case "redis":
		{
			cnf := cstr.Redis{}
			conf.Struct(&cnf, parent)
			out = engine.NewRedis(cnf.Host, cnf.Port, cnf.Db)
		}

	default:
		panik.Do("Unknown cache engine: %s", engn)
	}

	return out
}
