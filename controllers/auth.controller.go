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
	ctx.ShouldBind(&req)

	if req.Email == "" || req.Password == "" || req.ConfirmPassword == "" {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Email, password, and confirm password are required",
		})
		return
	}

	if req.Password != req.ConfirmPassword {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Password and confirm password do not match",
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
