package main

import (
	"golang-practice/user"
	// "net/http"
	// "github.com/gin-gonic/gin"
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

	userRepository := user.NewRepository(db)
	user := user.User{
		Name: "Test Aja",
	}

	userRepository.Save(user)
}