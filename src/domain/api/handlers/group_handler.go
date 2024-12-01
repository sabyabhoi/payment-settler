package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"sabyabhoi.com/payment-settler/src/domain/models"
	"sabyabhoi.com/payment-settler/src/domain/services"
)

type GroupHandler struct {
	groupService *services.GroupService
}

func NewGroupHandler(service *services.GroupService) *GroupHandler {
	return &GroupHandler{groupService: service}
}

func (h *GroupHandler) GetAllGroups(c *gin.Context) {
	groups, err := h.groupService.GetAllGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, groups)
}

func (h *GroupHandler) GetAllUsersInGroup(c *gin.Context) {
	groupId, err := strconv.Atoi(c.Query("groupId"))
  if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid group ID provided"})
		return
  }

  users, err := h.groupService.GetAllUsersInGroup(groupId)
  if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
  }
  c.JSON(http.StatusOK, users)
}

func (h *GroupHandler) GetGroupById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Make sure to pass in an integer as the ID"})
		return
	}

	group, err := h.groupService.GetGroupById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}
	c.JSON(http.StatusOK, group)
}

func (h *GroupHandler) CreateGroup(c *gin.Context) {
	var group models.Group

	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

  if group.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Group name cannot be empty."})
		return
  }

	uid, err := h.groupService.CreateGroup(&group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Group successfully created",
		"group_id": uid,
	})
}
