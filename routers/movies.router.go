package routers

import (
	"be-cinevo/controllers"
	"github.com/gin-gonic/gin"
)

func movieRouters(r *gin.RouterGroup){
	r.GET("", controllers.GetAllMovies)
	r.GET("/:id", controllers.GetMovieDetail)
	r.GET("/now-showing", controllers.GetNowShowingMovies)
	r.GET("/upcoming", controllers.GetUpComingMovies)
	r.GET("/genres", controllers.GetAllGenres )
	r.GET("/casts", controllers.GetAllCasts )
	r.GET("/directors", controllers.GetAllDirectors )
}