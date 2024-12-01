package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"sabyabhoi.com/payment-settler/src/domain/models"
	"sabyabhoi.com/payment-settler/src/domain/services"
)

type UserHandler struct {
	userService  *services.UserService
	groupService *services.GroupService
}

func NewUserHandler(userService *services.UserService, groupService *services.GroupService) *UserHandler {
	return &UserHandler{userService, groupService}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
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

func (h *UserHandler) GetAllGroupsForUserId(c *gin.Context) {
	uid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Make sure to pass in correct user id"})
		return
	}

  groups, err := h.userService.GetAllGroupsForUserId(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch groups for user"})
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (h *UserHandler) AddGroupToUser(c *gin.Context) {
	uid, err := strconv.Atoi(c.Query("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Make sure to pass in correct user id"})
		return
	}

	gid, err := strconv.Atoi(c.Query("groupid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Make sure to pass in correct group id"})
		return
	}

	user, err := h.userService.GetUserById(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group, err := h.groupService.GetGroupById(gid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.userService.AddGroupToUser(user, group)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Group successfully added for user"})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect email format specified"})
		return
	}

	uid, err := h.userService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User successfully created",
		"user_id": uid,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := h.userService.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponse)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.userService.DeleteUser(uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func isValidEmail(email string) bool {
	regex := `[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`

	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
