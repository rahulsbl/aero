package orm

import (
	"github.com/jinzhu/gorm"
	"github.com/thejackrabbit/aero/db/cstr"
)

var engines map[string]gorm.DB

func init() {
	engines = make(map[string]gorm.DB)

	// default to SingularTable
	DoOrmInit(func(o *gorm.DB) {
		o.SingularTable(true)
	})
}

func Get(useMaster bool) gorm.DB {
	s := cstr.Get(useMaster)
	return From(s.Engine, s.Conn)
}

func ReadConfig(container string) gorm.DB {
	s := cstr.ReadConfig(container)
	return From(s.Engine, s.Conn)
}

func From(engine string, conn string) gorm.DB {
	var ormObj gorm.DB
	var ok bool
	var err error

	if ormObj, ok = engines[conn]; ok {
		return ormObj
	}
	// http://go-database-sql.org/accessing.html
	// the sql.DB object is designed to be long-lived
	if ormObj, err = gorm.Open(engine, conn); err == nil {
		if ormInit != nil {
			for _, fn := range ormInit {
				fn(&ormObj)
			}
		}
		engines[conn] = ormObj
		return engines[conn]
	} else {
		panic(err)
	}
}

// orm initializers
var ormInit []func(*gorm.DB) = make([]func(*gorm.DB), 0)

func DoOrmInit(fn func(*gorm.DB)) {
	// TODO: use mutex
	ormInit = append(ormInit, fn)
}
