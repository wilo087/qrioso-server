package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
