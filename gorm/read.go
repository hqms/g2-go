package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed connect to DB")
	}
	var user User

	db.First(&user)
	fmt.Println(user.Name)

	result := map[string]interface{}{}
	db.Model(&User{}).First(result)
	fmt.Println(result)

	result = map[string]interface{}{}
	db.Table("users").Take(&result)
	fmt.Println(result)
}
