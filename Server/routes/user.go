package routes

import (
	"assessment-bootcamp/Server/auth"
	"assessment-bootcamp/Server/config"
	"assessment-bootcamp/Server/handler"
	"assessment-bootcamp/Server/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB = config.Connect()
	userRepository          = user.NewRepository(DB)
	userService             = user.UserNewService(userRepository)
	authService             = auth.NewService()
	userHandler             = handler.NewUserHandler(userService, authService)
)

func UserRoute(r *gin.Engine) {
	r.POST("/users/register", userHandler.CreateUserHandler)
	r.POST("/users/login", userHandler.LoginUserHandler)
	r.PUT("/users/:id", handler.Middleware(userService, authService), userHandler.UpdateUserByIDHandler)
	r.GET("/users/:id", handler.Middleware(userService, authService), userHandler.ShowUserByIdHandler)
}
