package main

import (
	"golang-practice/user"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main()  {
	dsn := "root:@tcp(127.0.0.1:3306)/golang_practice?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil{
		log.Fatal(err.Error())
	}

	fmt.Println("Database Connected")

	var users []user.User
	length := len(users)

	fmt.Println(length)
	
	//pointer type
	db.Find(&users)

	length = len(users)
	fmt.Println(length)

	for _, user := range users{
		fmt.Println(user.Name)
		fmt.Println(user.Email)
		fmt.Println("========")
	}

}