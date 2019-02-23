package main

import (
	"fmt"
	"github.com/lotteryjs/winning/log"
	"math/rand"
	"os"
	"time"
)

// Logger
var logger *log.Logger

// The only one init function.
func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	log.SetLevel("warn")
	logger = log.NewLogger(os.Stdout)
}

// Entry point.
func main() {
	fmt.Println("I AM WINNING")
}
