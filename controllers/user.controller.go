package controllers

import (
	"be-cinevo/dto"
	"be-cinevo/models"
	"be-cinevo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized!",
		})
		return
	}

	user, err := models.FindUserById(userId.(int))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to find user",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Get user info success!",
		Results: user,
	})
}

func UpdateUserInfo(ctx *gin.Context) {
	userId, exists := ctx.Get("userId")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized!",
		})
		return
	}

	req := dto.UpdatedUser{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid Format Request!",
		})
		return
	}

	err = models.GetUpdatedUserInfo(userId.(int), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to find user",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "User Info Updated Successfully!",
	})
}
