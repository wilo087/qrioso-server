package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultHandler struct {
	user *UserHandler
}

func NewDefaultHandler() *DefaultHandler {
	return &DefaultHandler{user: NewUserHandler()}
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

	v1.GET("/users", d.user.Get)
	v1.POST("/users", d.user.Create)

	return r
}
