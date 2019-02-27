package config

import (
	"github.com/lotteryjs/configor"
)

// Configuration is stuff that can be configured externally per env variables or config file (config.yml).
type Configuration struct {
	Server struct {
		Port            int `default:"80"`
		ResponseHeaders map[string]string
	}
	Database struct {
		Dialect    string `default:"sqlite3"`
		Connection string `default:"data/gotify.db"`
	}
	DefaultUser struct {
		Name string `default:"admin"`
		Pass string `default:"admin"`
	}
	PassStrength int `default:"10"`
}

// Get returns the configuration extracted from env variables or config file.
func Get() *Configuration {
	conf := new(Configuration)
	err := configor.New(&configor.Config{EnvironmentPrefix: "WINNING"}).Load(conf, "config.yml", "/etc/winning/config.yml")
	if err != nil {
		panic(err)
	}
	return conf
}
