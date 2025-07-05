package routers

import (
	"be-cinevo/controllers"
	"github.com/gin-gonic/gin"
)

func movieRouters(r *gin.RouterGroup){
	r.GET("", controllers.GetAllMovies)
	r.GET("/now-showing", controllers.GetNowShowingMovies)
	r.GET("/upcoming", controllers.GetUpComingMovies)
}