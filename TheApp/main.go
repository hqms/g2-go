package main

import (
	"github.com/getsentry/sentry-go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"theapp/Config"
	"theapp/Models"
	"theapp/Routes"
	"time"
)


func main() {
	var err error
	Config.DB, err = gorm.Open(sqlite.Open(Config.DbUrl(Config.BuildDBConfig())), &gorm.Config{})
	if err != nil{
		panic("Cannot connect to DB")
	}

	//defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	defer sentry.Flush(2 * time.Second)
	
	r := Routes.SetupRouter()
	r.Run()
}
