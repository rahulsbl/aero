package cstr

import (
	"fmt"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rahulsbl/aero/conf"

)

func init() {
	initAllMasterAndSlaves();
}

type Storage struct {
	Engine string
	Conn   string
	Mdb    string // mongo database name
}

var dbMaster = make(map[string]Storage, 0)
var dbSlaves = make(map[string][]Storage, 0)
var master Storage
var slaves []Storage = make([]Storage, 0)


func initAllMasterAndSlaves() {

	// Master
	lookup := "databases"
	if conf.Exists(lookup) {
		readDbConfigs := conf.Get("", lookup);
		for k := range readDbConfigs.(map[string]interface{}) {
			var slavesDb []Storage = make([]Storage, 0)
		 	masterLookup := lookup + "." + k + ".master"
		 	dbMaster[k] = ReadConfig(masterLookup)
		 	slaveLookup := lookup + "." + k + ".slaves";
		 	slaves := conf.StringSlice([]string{}, slaveLookup)
			for _, container := range slaves {
				slavesDb = append(slavesDb, ReadConfig(container))
			}
			dbSlaves[k] = slavesDb
		}
	}
}

func initMasterAndSlaves() {

	// Master
	lookup := "database.master"
	if conf.Exists(lookup) {
		master = ReadConfig(lookup)
	}

	// Slaves
	lookup = "database.slaves"
	if conf.Exists(lookup) {
		slav := conf.StringSlice([]string{}, lookup)
		for _, container := range slav {
			slaves = append(slaves, ReadConfig(container))
		}
	}
}

func ReadConfig(container string) (s Storage) {

	// Get the "type" of the db
	if !conf.Exists(container) {
		panic(fmt.Sprintf("Configuration under %s not found", container))
	}

	s.Engine = conf.String("", container, "engine")

	switch s.Engine {
	case "mysql", "maria", "mariadb":
		{
			m := Mysql{}
			conf.Struct(&m, container)

			s.Engine = "mysql"
			s.Conn = m.Cstr()
		}

	case "postgres":
		{
			username := conf.String("", container, "username")
			password := conf.String("", container, "password")
			host := conf.String("", container, "host")
			port := conf.String("", container, "port")
			db := conf.String("", container, "db")
			sslmode := conf.String("disable", container, "sslmode")

			auth := ""
			if username != "" || password != "" {
				auth = username + ":" + password + "@"
			}

			s.Engine = "postgres"
			s.Conn = fmt.Sprintf("postgres://%s%s:%s/%s?sslmode=%s",
				auth,
				host, port, db, sslmode,
			)
		}

	case "sqlite3":
		{
			q := Sqlite{}
			conf.Struct(&q, container)

			s.Engine = "sqlite3"
			s.Conn = q.Cstr()
		}

	case "mongo", "mongodb":
		{
			m := Mongodb{}
			conf.Struct(&m, container)

			s.Engine = "mongodb"
			s.Conn = m.Cstr()
			s.Mdb = m.Db
		}

	case "memcache":
		{
			m := Memcache{}
			conf.Struct(&m, container)

			s.Engine = "memcache"
			s.Conn = m.Cstr()
		}

	default:
		panic(fmt.Sprintf("Unsupported db %s", s.Engine))
	}

	return
}

func Get(useMaster bool) Storage {
	if useMaster {
		return master
	}

	if len(slaves) == 0 {
		return master
	}

	return slaves[rand.Intn(len(slaves))]
}

func SelectConfig(dbName string, useMaster bool) Storage {
	
	if useMaster {
		db, ok := dbMaster[dbName]
		if !ok || db == (Storage{}) {         
            panic(fmt.Sprintf("Unsupported db %s", dbName))
        }
		return dbMaster[dbName]
	}

	if len(dbSlaves[dbName]) == 0 {
		return dbMaster[dbName]
	}
	return dbSlaves[dbName][rand.Intn(len(dbSlaves[dbName]))]
}
