package que

import (
	"strings"
	"time"

	"github.com/thejackrabbit/aero/conf"
	"github.com/thejackrabbit/aero/db/cstr"
	"github.com/thejackrabbit/aero/engine"
	"github.com/thejackrabbit/aero/panik"
)

type Queue interface {
	Push(data []byte) error
	Pop() ([]byte, error)
	PopWait(time.Duration) ([]byte, error)
	Len() (int, error)
	Close()
}

func NewQueue(container ...string) (out Queue) {
	parent := strings.Join(container, ".")

	engn := conf.String("", parent, "engine")
	panik.If(engn == "", "queue engine is not specified under %s", parent)

	switch engn {
	case "redis":
		{
			cnf := cstr.Redis{}
			conf.Struct(&cnf, parent)
			panik.If(cnf.Name == "", "Redis queue name missing")
			out = engine.NewRedis2(cnf.Host, cnf.Port, cnf.Db, cnf.Name)
		}

	default:
		panik.Do("Unknown queue engine: %s", engn)
	}

	return out
}
