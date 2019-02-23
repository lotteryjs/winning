package model

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/lotteryjs/winning/log"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// Logger
var logger = log.NewLogger(os.Stdout)

// Version of Winning.
const Version = "1.0.0"

// Conf of Winning.
var Conf *Configuration

// Table prefix.
const tablePrefix = "wnn_"

// Configuration (winning.json).
type Configuration struct {
	Server        string // server scheme, host and port
	LogLevel      string // logging level: trace/debug/info/warn/error/fatal
	ShowSQL       bool   // whether print sql in log
	SessionSecret string // HTTP session secret
	SessionMaxAge int    // HTTP session max age (in seciond)
	RuntimeMode   string // runtime mode (dev/prod)
	MySQL         string // MySQL connection URL
	Port          string // listen port
}

// LoadConf loads the configurations. Command-line arguments will override configuration file.
func LoadConf() {
	version := flag.Bool("version", false, "prints current winning version")
	confPath := flag.String("conf", "winning.json", "path of winning.json")
	confServer := flag.String("server", "", "this will override Conf.Server if specified")
	confLogLevel := flag.String("log_level", "", "this will override Conf.LogLevel if specified")
	confShowSQL := flag.Bool("show_sql", false, "this will override Conf.ShowSQL if specified")
	confRuntimeMode := flag.String("runtime_mode", "", "this will override Conf.RuntimeMode if specified")
	confMySQL := flag.String("mysql", "", "this will override Conf.MySQL if specified")
	confPort := flag.String("port", "", "this will override Conf.Port if specified")

	flag.Parse()

	if *version {
		fmt.Println(Version)

		os.Exit(0)
	}

	bytes, err := ioutil.ReadFile(*confPath)
	if nil != err {
		logger.Fatal("loads configuration file [" + *confPath + "] failed: " + err.Error())
	}

	Conf = &Configuration{}
	if err = json.Unmarshal(bytes, Conf); nil != err {
		logger.Fatal("parses [winning.json] failed: ", err)
	}

	log.SetLevel(Conf.LogLevel)
	if "" != *confLogLevel {
		Conf.LogLevel = *confLogLevel
		log.SetLevel(*confLogLevel)
	}

	if *confShowSQL {
		Conf.ShowSQL = true
	}

	if "" != *confRuntimeMode {
		Conf.RuntimeMode = *confRuntimeMode
	}

	if "" != *confServer {
		Conf.Server = *confServer
	}

	time := strconv.FormatInt(time.Now().UnixNano(), 10)
	logger.Debugf("${time} [%s]", time)

	if "" != *confMySQL {
		Conf.MySQL = *confMySQL
	}

	if "" != *confPort {
		Conf.Port = *confPort
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	logger.Debugf("configurations [%#v]", Conf)
}
