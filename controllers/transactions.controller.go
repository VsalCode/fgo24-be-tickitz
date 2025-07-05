package controllers

import (
	"be-cinevo/models"
	"be-cinevo/utils"
	"net/http"
	"github.com/gin-gonic/gin"
);

// BookingTickets godoc
// @Summary Booking Tickets
// @Description add booking movie ticket transactions
// @Tags Transactions
// @Produce json
// @Param ticket body models.Transactions true "Data Transactions"
// @Success 201 {object} utils.Response
// @Security BearerAuth
// @Router /transactions [post]
func BookingTicket(ctx *gin.Context){
	userId, exists := ctx.Get("userId")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized!",
		})
		return
	}

	
	req := models.Transactions{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid Request!",
			Errors: err.Error(),
		})
		return
	}

	err = models.HandleBookingTicket(userId.(int), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to proccess transaction!",
			Errors: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "Booking Ticket Successfully!",
	})
}

// TransactionsHistory godoc
// @Summary Transactions History
// @Description Get User Transactions History
// @Tags Transactions
// @Produce json
// @Success 200 {object} utils.Response
// @Security BearerAuth
// @Router /transactions [get]
func HistoryTransactions(ctx *gin.Context){
		userId, exists := ctx.Get("userId")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized!",
		})
		return
	}

	historys, err := models.FindHistoryByUserId(userId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed To Get History Transactions!",
			Errors: err.Error(),
		})
		return
	} 

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Success to Get Transactions History!",
		Results: historys,
	})
}