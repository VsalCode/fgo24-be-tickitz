package controllers

import (
	"be-cinevo/dto"
	"be-cinevo/models"
	"be-cinevo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserInfo godoc
// @Summary Get User Profile
// @Description Get the profile information of current logged-in user
// @Tags User
// @Produce json
// @Success 200 {object} utils.Response{results=models.User}
// @Security BearerAuth
// @Router /user [get]
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

// UpdateUserInfo godoc
// @Summary Update User Profile
// @Description Update the profile information of the logged-in user
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.UpdatedUser true "User data to update"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security BearerAuth
// @Router /user [patch]
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
