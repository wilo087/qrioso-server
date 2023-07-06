package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wilo0087/qrioso-server/internal/service"
)

type DefaultHandler struct {
	user *UserHandler
}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{user: NewUserHandler(&service.UserService{})}
}

func StartApiRoutes(d *DefaultHandler) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// r.POST("/inform", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"status": "ok"})
	// })

	v1.GET("/users/:id", d.user.GetUserByID)
	v1.POST("/users", d.user.Create)

	return r
}
