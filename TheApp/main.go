package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"theapp/Config"
	"theapp/Models"
	"theapp/Routes"
)

func main() {
	var err error
	Config.DB, err = gorm.Open(sqlite.Open(Config.DbUrl(Config.BuildDBConfig())), &gorm.Config{})
	if err != nil{
		panic("Cannot connect to DB")
	}

	//defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()

	r.Run()
}
