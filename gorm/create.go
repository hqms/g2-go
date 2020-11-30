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
	Address []Address
}

type Address struct {
	gorm.Model
	Alamat string
	UserID uint
}


func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to DB")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Address{})

	joni := User{Name: "HQM with address", Age: 17, Address: []Address{{ Alamat: "Disini"}, {Alamat: "dimana mana"}}}
	result := db.Create(&joni)

	fmt.Println("ID: ", joni.ID)
	fmt.Println("If error return :", result.Error)
	fmt.Println("Row retruned ", result.RowsAffected)

}
