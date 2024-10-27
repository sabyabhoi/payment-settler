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

	repo := repositories.NewUserRepository(db)
	service := services.NewUserService(repo)
	handler := handlers.NewUserHandler(service)
	router := routes.NewRouter(handler)

	r := router.SetupRoutes()

	r.Run(":8080")
}
