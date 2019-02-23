package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/lotteryjs/winning/log"
	"github.com/lotteryjs/winning/model"
	"github.com/lotteryjs/winning/service"
)

// Logger
var logger *log.Logger

// The only one init function.
func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	log.SetLevel("warn")
	logger = log.NewLogger(os.Stdout)

	model.LoadConf()

	if "dev" == model.Conf.RuntimeMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
}

// Entry point.
func main() {
	service.ConnectDB()
	fmt.Println("I AM WINNING")
}
