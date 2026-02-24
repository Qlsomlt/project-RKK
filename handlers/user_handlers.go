package handlers

import (
	"kode/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{s}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	role := c.GetString("role")

	users, err := h.service.GetAllUsers(role)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
