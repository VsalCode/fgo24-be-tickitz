package routers

import (
	"be-cinevo/controllers"
	"be-cinevo/middlewares"
	"github.com/gin-gonic/gin"
)

func userRouters(r *gin.RouterGroup){
	r.Use(middlewares.VerifyToken())
	r.GET("", controllers.GetUserInfo )
	r.PATCH("", controllers.UpdateUserInfo )
}