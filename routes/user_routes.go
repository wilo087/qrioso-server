package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wilo0087/qrioso-server/internal/handlers"
)

func UserRoutes(r *gin.Engine, uc *handlers.UserHandler) {
	userRoutes := r.Group("/v1/users")

	userRoutes.GET("/:id", uc.GetUserByID)
}
