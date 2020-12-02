package main

import (
	"github.com/getsentry/sentry-go"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
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

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://52b4560b8c1046a8abdbafbdad51d369@o485449.ingest.sentry.io/5540851",
	});  err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	defer sentry.Flush(2 * time.Second)
	
	r := Routes.SetupRouter()

	r.Run()
}
