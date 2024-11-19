package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sabyabhoi.com/payment-settler/src/domain/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Make sure to pass in an integer as the ID"})
		return
	}

	user, err := h.userService.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")

	user, err := h.userService.GetUserByEmail(email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
