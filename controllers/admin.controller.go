package controllers

import (
	"be-cinevo/dto"
	"be-cinevo/models"
	"be-cinevo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddMovie godoc
// @Summary Add a new movie
// @Description Add a new movie (admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Param movie body dto.MovieRequest true "Movie data"
// @Success 201 {object} utils.Response
// @Security BearerAuth
// @Router /admin [post]
func AddMovie(ctx *gin.Context) {
	role, _ := ctx.Get("role")
	if role != "admin" {
		ctx.JSON(http.StatusForbidden, utils.Response{
			Success: false,
			Message: "Forbidden: Only admin can add movies",
		})
		return
	}

	var req dto.MovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request format",
			Errors:  err.Error(),
		})
		return
	}

	err := models.CreateNewMovie(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to add movie",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "Movie added successfully",
	})
}

func DeleteMovie(ctx *gin.Context) {
	role, _ := ctx.Get("role")
	if role != "admin" {
		ctx.JSON(http.StatusForbidden, utils.Response{
			Success: false,
			Message: "Forbidden: Only admin can add movies",
		})
		return
	}

	id := ctx.Param("id")
	IdInt, _ := strconv.Atoi(id)

	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "User ID is required",
		})
		return
	}

	err := models.DeleteMovieById(IdInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to Delete Movie",
			Results: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Delete Movie Successfully!",
	})
	return
}
