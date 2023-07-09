package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wilo0087/qrioso-server/internal/handlers"
	"github.com/wilo0087/qrioso-server/internal/repository"
	"github.com/wilo0087/qrioso-server/internal/service"
	"gorm.io/gorm"
)

type UserRoutes struct {
	userHandler *handlers.UserHandler
}

func NewUserRoutes(db *gorm.DB) *UserRoutes {

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	return &UserRoutes{userHandler: userHandler}
}

func RegisterUserRoutes(ur *UserRoutes, r *gin.Engine) {
	userRoutes := r.Group("/v1/users")
	userRoutes.GET("/:id", ur.userHandler.GetUserByID)
}
