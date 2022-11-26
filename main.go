package main

import (
	"log"

	"github.com/amarmaulana95/backend-portal/auth"
	"github.com/amarmaulana95/backend-portal/handler"
	user "github.com/amarmaulana95/backend-portal/users"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root@tcp(127.0.0.1:3306)/portal?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()       //membuat routing
	api := router.Group("api/v1") //api group versioning (untuk dkebutuhan aja)

	api.POST("/users", userHandler.RegisterUser) //endpoint
	api.POST("/sessions", userHandler.Login)     //login

	router.Run()
}
