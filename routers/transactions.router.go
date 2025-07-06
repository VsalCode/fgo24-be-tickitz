package routers

import (
	"be-cinevo/controllers"
	"be-cinevo/middlewares"

	"github.com/gin-gonic/gin"
)

func transactionRouters(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.POST("", controllers.BookingTicket)
	r.GET("", controllers.TicketResult)
	r.GET("/history", controllers.HistoryTransactions)
}
