package handlers

import (
	"kode/services"
	"kode/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service services.UserService
}

func NewAuthHandler(s services.UserService) *AuthHandler {
	return &AuthHandler{s}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	c.BindJSON(&body)

	err := h.service.Register(body.Name, body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registered"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	c.BindJSON(&body)

	token, err := h.service.Login(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) Logout(c *gin.Context) {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utils.Error(c, 401, "authorization required")
		return
	}

	err := h.authService.Logout(authHeader)
	if err != nil {
		utils.Error(c, 400, "logout failed")
		return
	}

	utils.OK(c, "logged out successfully")
}
