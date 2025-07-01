package controllers

import (
	"be-cinevo/dto"
	"be-cinevo/models"
	"be-cinevo/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {

	req := dto.RegisterRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request format",
			Errors:  err.Error(),
		})
		return
	}

	result, err := models.GetNewUser(req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to register user",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User registered successfully",
		Results: result,
	})

}

func LoginUser() {}
