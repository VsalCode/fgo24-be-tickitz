package routers

import (
	"be-cinevo/controllers"

	"github.com/gin-gonic/gin"
)

func movieRouters(r *gin.RouterGroup){
	r.GET("/now-showing", controllers.NowShowingMovies)
	r.GET("/upcoming", controllers.UpComingMovies)
}