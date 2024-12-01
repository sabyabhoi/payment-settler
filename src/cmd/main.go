package main

import (
	"fmt"

	"sabyabhoi.com/payment-settler/src/domain/api/handlers"
	"sabyabhoi.com/payment-settler/src/domain/api/routes"
	"sabyabhoi.com/payment-settler/src/domain/config"
	"sabyabhoi.com/payment-settler/src/domain/services"
	"sabyabhoi.com/payment-settler/src/infrastructure/database"
	"sabyabhoi.com/payment-settler/src/infrastructure/repositories"
)

func main() {

	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Failed to read config")
		return
	}

	db, err := database.NewConnection(config.ToDatabaseConfig())
	if err != nil {
		fmt.Println("Failed to load database")
		return
	}

	userRepo := repositories.NewUserRepository(db)
	groupRepo := repositories.NewGroupRepository(db)
  transactionRepo := repositories.NewTransactionRepository(db)

	userService := services.NewUserService(userRepo)
	groupService := services.NewGroupService(groupRepo)
	transactionService := services.NewTransactionService(transactionRepo)

	userHandler := handlers.NewUserHandler(userService, groupService)
  groupHandler := handlers.NewGroupHandler(groupService)
  transactionHandler := handlers.NewTransactionHandler(transactionService)

	router := routes.NewRouter(userHandler, groupHandler, transactionHandler)

	r := router.SetupRoutes()

	r.Run(":8080")
}
