package cstr

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/thejackrabbit/aero/conf"
	"github.com/thejackrabbit/aero/panik"
	"math/rand"
	"net/url"
)

func init() {
	initMasterAndSlaves()
}

type Schema struct {
	Storage string
	Conn    string
	Mdb     string // mongo database name
}

var master Schema
var slaves []Schema = make([]Schema, 0)

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

func ReadConfig(container string) (s Schema) {

	// Get the "type" of the db
	if !conf.Exists(container) {
		panik.Do("Configuration under %s not found", container)
	}

	s.Storage = conf.String("", container, "storage")

	switch s.Storage {
	case "mysql", "maria", "mariadb":
		{

			username := conf.String("", container, "username")
			password := conf.String("", container, "password")
			host := conf.String("", container, "host")
			port := conf.String("", container, "port")
			db := conf.String("", container, "db")
			timezone := conf.String("", container, "timezone")

			s.Storage = "mysql"
			s.Conn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s",
				username, password,
				host, port, db,
				url.QueryEscape(timezone),
			)
		}

	case "sqlite3":
		{
			path := conf.String("", container, "path")

			s.Storage = "sqlite3"
			s.Conn = path
		}

	case "mongo", "mongodb":
		{
			username := conf.String("", container, "username")
			password := conf.String("", container, "password")
			host := conf.String("", container, "host")
			port := conf.String("", container, "port")
			db := conf.String("", container, "db")
			replicas := conf.String("", container, "replicas")
			options := conf.String("", container, "options")

			if port != "" {
				port = ":" + port
			}
			if replicas != "" {
				replicas = "," + replicas
			}
			if options != "" {
				options = "?" + options
			}
			auth := ""
			if username != "" || password != "" {
				auth = username + ":" + password + "@"
			}

			s.Storage = "mongo"
			s.Conn = fmt.Sprintf("mongodb://%s%s%s%s/%s%s",
				auth, host, port, replicas,
				db, options,
			)
			s.Mdb = db
		}
	default:
		panik.Do("Unsupported db %s", s.Storage)
	}

	return
}

func Get(useMaster bool) Schema {
	if useMaster {
		return master
	}

	if len(slaves) == 0 {
		return master
	}

	return slaves[rand.Intn(len(slaves))]
}
