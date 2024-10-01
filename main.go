package main

import (
	"github.com/HIUNCY/rest-api-go/handler"
	"github.com/HIUNCY/rest-api-go/model"
	"github.com/HIUNCY/rest-api-go/repository"
	"github.com/HIUNCY/rest-api-go/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("sql12734405:HCwYzkUDCB@tcp(sql12.freesqldatabase.com:3306)/sql12734405?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})

	// USER
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/login", userHandler.Login)
		user.POST("/register", userHandler.Register)
	}
	r.Run()
}
