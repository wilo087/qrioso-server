package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wilo0087/qrioso-server/internal/model/dto"
	"github.com/wilo0087/qrioso-server/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(u *service.UserService) *UserHandler {
	return &UserHandler{userService: u}
}

func (u *UserHandler) GetUserByID(c *gin.Context) {
	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		fmt.Println("Error al analizar el UUID:", err)
		return
	}

	user, err := u.userService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "404"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "data": user})
}

func (u *UserHandler) GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (u *UserHandler) Create(c *gin.Context) {
	var user dto.CreateUserDTO

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "ERROR"})
		return
	}

	log.Println(user)
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
