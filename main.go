package main

import (
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/lotteryjs/winning/controller"
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
	router := controller.MapRoutes()
	server := &http.Server{
		Addr:    "0.0.0.0:" + model.Conf.Port,
		Handler: router,
	}

	handleSignal(server)

	logger.Infof("Winning (v%s) is running [%s]", model.Version, model.Conf.Server)
	server.ListenAndServe()
}

// handleSignal handles system signal for graceful shutdown.
func handleSignal(server *http.Server) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		logger.Infof("got signal [%s], exiting winning now", s)
		if err := server.Close(); nil != err {
			logger.Errorf("server close failed: " + err.Error())
		}

		service.DisconnectDB()

		logger.Infof("Winning exited")
		os.Exit(0)
	}()
}
