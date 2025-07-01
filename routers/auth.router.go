package routers;

import (
	"github.com/gin-gonic/gin"
	"be-cinevo/controllers"
)

func authRouters(r *gin.RouterGroup){
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser )
} 