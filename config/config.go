package config

import (
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

var conf = main()

// PORT ...
const PORT = "8000"

// HOST ...
var HOST = conf.HOST + ":"

// URL ...
var URL = conf.URL

// SiteTitle ...
var SiteTitle = conf.SiteTitle

// HashSalt ...
var HashSalt = conf.HashSalt

// DbType ...
const DbType = "sqlite3"

// DbURL ...
var DbURL = conf.DbDatabase

// JwtExp ...
var JwtExp = time.Now().Add(time.Hour * 168).Unix()

// GOPATH ...
var GOPATH = os.Getenv("GOPATH")

// ProjectPath ...
var ProjectPath = GOPATH + "/src/go-echo-vue/"

type (
	// ENV ...
	ENV struct {
		Production Config
		Develop    Config
		Test       Config
	}

	// Config ...
	Config struct {
		DbDatabase string `toml:"DB_DATABASE"`
		DbHost     string `toml:"DB_HOST"`
		DbUser     string `toml:"DB_USER"`
		DbPass     string `toml:"DB_PASS"`
		HOST       string `toml:"HOST"`
		URL        string `toml:"URL"`
		SiteTitle  string `toml:"SITE_TITLE"`
		HashSalt   string `toml:"HASH_SALT"`
	}
)

func main() Config {
	env := ENV{}
	confPath := ProjectPath + "config/config.tml"
	_, err := toml.DecodeFile(confPath, &env)

	if err != nil {
		panic(err)
	}

	res := Config{}
	osEnv := os.Getenv("DOCUMENT_ENV")
	if osEnv == "production" {
		res = env.Production
	} else if osEnv == "test" {
		res = env.Test
	} else {
		res = env.Develop
	}
	return res
}
