package routers

import (
	"be-cinevo/controllers"
	"be-cinevo/middlewares"

	"github.com/gin-gonic/gin"
)

func authRouters(r *gin.RouterGroup) {
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.POST("/forgot-password", controllers.ForgotPassword)
	r.POST("/reset-password", controllers.ResetPassword)
	r.POST("/logout", middlewares.VerifyToken(), controllers.LogoutUser)
}
