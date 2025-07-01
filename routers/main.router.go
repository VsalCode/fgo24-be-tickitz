package routers;

import (
	"github.com/gin-gonic/gin"
)

func CombineRouters(r *gin.Engine){
	authRouters(r.Group(""))
}