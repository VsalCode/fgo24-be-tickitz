package controllers

import (
	"net/http"
	"be-cinevo/utils"
	"be-cinevo/models"
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

