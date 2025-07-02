package routers

import (
	"be-cinevo/controllers"
	"be-cinevo/middlewares"
	"github.com/gin-gonic/gin"
)

func adminRouters(r *gin.RouterGroup){
	r.Use(middlewares.VerifyToken())
	r.POST("", controllers.AddMovie )
}