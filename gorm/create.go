package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name    string
	Age     int
	Address string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to DB")
	}

	db.AutoMigrate(&User{})

	joni := User{Name: "HQM", Age: 17, Address: "disitu"}
	result := db.Create(&joni)

	fmt.Println("ID: ", joni.ID)
	fmt.Println("If error return :", result.Error)
	fmt.Println("Row retruned ", result.RowsAffected)
}
