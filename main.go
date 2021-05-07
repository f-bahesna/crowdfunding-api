package main

import (
	"golang-practice/user"
	"golang-practice/handler"
	// "net/http"
	"github.com/gin-gonic/gin"
	"log"
	// "fmt"

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
	userService := user.NewService(userRepository)

	// userService.SaveAvatar(1, "image/1-profile.png")

	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.EmailAvailability)

	router.Run()
}