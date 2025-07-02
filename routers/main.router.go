package routers

import (
	"github.com/gin-gonic/gin"

	"be-cinevo/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CombineRouters(r *gin.Engine) {
	authRouters(r.Group("/auth"))
	movieRouters(r.Group("/movies"))
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
