package conf

import (
	"fmt"
	"github.com/jacobstr/confer"
	"path/filepath"
)

// TODO:
// - allow command line arguments to be passed
// - eg. --APP_PORT=1234, should be bubbled to the right place

var configuration *confer.Config

func init() {
	loadDefaultConfig()
}

func loadDefaultConfig() {
	configuration = confer.NewConfig()
	seek := []string{
		// least priority
		"./config/config.yaml",
		"config.yaml",

		// next priority
		"./config/dev.yaml",
		"dev.yaml",

		// highest priority
		"./config/production.yaml",
		"production.yaml",
	}
	var err error
	var files []string = make([]string, 0)
	for _, f := range seek {
		tmp := confer.NewConfig()
		err = tmp.ReadPaths(f)
		if err == nil {
			abs, _ := filepath.Abs(f)
			files = append(files, abs)
		}
	}

	if len(files) == 0 {
		fmt.Println("No yaml configuration file found.")
	} else {
		configuration.ReadPaths(files...)
		fmt.Println("Loading configurations:", len(files), "file(s)")
		for i := 0; i < len(files)-1; i++ {
			fmt.Print(files[i], " → ")
		}
		fmt.Print(files[len(files)-1], "\n")
	}
}

func Get(key string, defValue interface{}) interface{} {
	if configuration.IsSet(key) {
		return configuration.Get(key)
	} else {
		return defValue
	}
}

func Int(key string, defValue int) int {
	if Exists(key) {
		return configuration.GetInt(key)
	} else {
		return defValue
	}
}

func String(key string, defValue string) string {
	if Exists(key) {
		return configuration.GetString(key)
	} else {
		return defValue
	}
}

func StringSlice(key string, defValue []string) []string {
	if Exists(key) {
		return configuration.GetStringSlice(key)
	} else {
		return defValue
	}
}

func Bool(key string, defValue bool) bool {
	if Exists(key) {
		return configuration.GetBool(key)
	} else {
		return defValue
	}
}

func Exists(key string) bool {
	return configuration.IsSet(key)
}
