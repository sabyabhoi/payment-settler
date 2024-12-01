package routes

import (
	"github.com/gin-gonic/gin"
	"sabyabhoi.com/payment-settler/src/domain/api/handlers"
)

type Router struct {
	userHandler  *handlers.UserHandler
	groupHandler *handlers.GroupHandler
}

func NewRouter(userHandler *handlers.UserHandler, groupHandler *handlers.GroupHandler) *Router {
	return &Router{userHandler, groupHandler}
}

func (r *Router) SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	userApi := router.Group("/users")
	{
		userApi.GET("/", r.userHandler.GetAllUsers)
		userApi.GET("/:id", r.userHandler.GetUserById)
		userApi.GET("/group", r.userHandler.GetAllGroupsForUserId)
		userApi.GET("/email/:email", r.userHandler.GetUserByEmail)

		userApi.POST("/", r.userHandler.CreateUser)
		userApi.POST("/group", r.userHandler.AddGroupToUser)

		userApi.DELETE("/:id", r.userHandler.DeleteUser)
	}

	groupApi := router.Group("/groups")
	{
		groupApi.GET("/", r.groupHandler.GetAllGroups)
		groupApi.GET("/:id", r.groupHandler.GetGroupById)
		groupApi.GET("/user", r.groupHandler.GetAllUsersInGroup)

		groupApi.POST("/", r.groupHandler.CreateGroup)
	}

	return router
}
