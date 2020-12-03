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

// @title TheApp API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
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
