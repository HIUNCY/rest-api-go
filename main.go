package main

import (
	"github.com/HIUNCY/rest-api-go/handler"
	"github.com/HIUNCY/rest-api-go/repository"
	"github.com/HIUNCY/rest-api-go/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:WyOuRPHnvSHOlRMmbRhzESwdWHCuPYVm@tcp(autorack.proxy.rlwy.net:56977)/tabungan?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// USER
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	// DEPOSIT
	depoRepository := repository.NewDepositRepository(db)
	depoService := service.NewDepositService(depoRepository)
	depoHandler := handler.NewDepositHandler(depoService)
	// TRANSACTION
	transactionRepository := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	r := gin.Default()
	r.Use(cors.Default())
	user := r.Group("/user")
	{
		user.GET("/list", userHandler.GetUserList)
		// user.GET("/balance/:nik", userHandler.GetUserList)
		user.POST("/login", userHandler.Login)
		user.POST("/register", userHandler.Register)
		user.PUT("/update", userHandler.Update)
		user.DELETE("/delete", userHandler.Delete)
	}
	depo := r.Group("/deposit")
	{
		depo.GET("/list", depoHandler.GetDepositList)
		depo.POST("/search", depoHandler.GetDepositByNik)
		depo.POST("/create", depoHandler.Create)
		depo.PUT("/update", depoHandler.Update)
		depo.DELETE("/delete", depoHandler.Delete)
	}
	transaction := r.Group("/transaction")
	{
		transaction.POST("/create", transactionHandler.CreateTransaction)
	}
	r.Run()
}
