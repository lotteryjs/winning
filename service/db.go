package service

import (
	"os"

	"github.com/lotteryjs/winning/log"
	"github.com/lotteryjs/winning/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql
)

// Logger
var logger = log.NewLogger(os.Stdout)

var db *gorm.DB

// ConnectDB connects to the database.
func ConnectDB() {
	var err error
	if "" != model.Conf.MySQL {
		db, err = gorm.Open("mysql", model.Conf.MySQL)
	} else {
		logger.Fatal("please specify database")
	}
	if nil != err {
		logger.Fatalf("opens database failed: " + err.Error())
	}
	logger.Debug("used [MySQL] as underlying database")

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
	db.LogMode(model.Conf.ShowSQL)
}

// DisconnectDB disconnects from the database.
func DisconnectDB() {
	if err := db.Close(); nil != err {
		logger.Errorf("Disconnect from database failed: " + err.Error())
	}
}

// Database returns the underlying database name.
func Database() string {
	return "MySQL"
}
