package route

import (
	"assessment-bootcamp/auth"
	"assessment-bootcamp/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB = config.Connection()
	authService          = auth.NewService()
	// userService          = user.NewRepository(DB)
)

func userRoute(r *gin.Engine) {
	r.POST("user/register")
}
