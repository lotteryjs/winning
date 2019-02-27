package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lotteryjs/winning/config"
	"github.com/lotteryjs/winning/mode"
	"github.com/lotteryjs/winning/model"
)

var (
	// Version the version of Gotify.
	Version = "unknown"
	// Commit the git commit hash of this version.
	Commit = "unknown"
	// BuildDate the date on which this binary was build.
	BuildDate = "unknown"
	// Mode the build mode
	Mode = mode.Dev
)

func main() {
	vInfo := &model.VersionInfo{Version: Version, Commit: Commit, BuildDate: BuildDate}
	mode.Set(Mode)

	fmt.Println("Starting Gotify version", vInfo.Version+"@"+BuildDate)
	rand.Seed(time.Now().UnixNano())
	conf := config.Get()

	db, err := database.New(conf.Database.Dialect, conf.Database.Connection, conf.DefaultUser.Name, conf.DefaultUser.Pass, conf.PassStrength, true)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
