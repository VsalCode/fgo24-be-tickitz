package routers

import (
	"be-cinevo/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CombineRouters(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://146.190.102.54:9402"}, 
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders: []string{"Authorization", "Origin", "Content-Type"},
	}))
	authRouters(r.Group("/auth"))
	movieRouters(r.Group("/movies"))
	adminRouters(r.Group("/admin"))
	userRouters(r.Group("/user"))
	transactionRouters(r.Group("/transactions"))
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
