package routes

import (
	"github.com/gin-gonic/gin"
	"sabyabhoi.com/payment-settler/src/domain/api/handlers"
)

type Router struct {
	userHandler *handlers.UserHandler
	// authHandler  *handlers.AuthHandler
	// authMiddleware *middlewares.AuthMiddleware
}

func NewRouter(userHandler *handlers.UserHandler) *Router {
	return &Router{userHandler}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/users")
	{
    api.GET("/", r.userHandler.GetUsers)
		api.GET("/:id", r.userHandler.GetUserById)
		api.GET("/email/:email", r.userHandler.GetUserByEmail)

    api.POST("/", r.userHandler.CreateUser)

    api.DELETE("/:id", r.userHandler.DeleteUser)
	}

	return router
}
