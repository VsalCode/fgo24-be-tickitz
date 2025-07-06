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

// DeleteMovie godoc
// @Summary Delete a movie
// @Description Delete movie by movie id (admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} utils.Response
// @Security BearerAuth
// @Router /admin/{id} [delete]
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
	idInt, _ := strconv.Atoi(id)

	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "User ID is required",
		})
		return
	}

	err := models.DeleteMovieById(idInt)
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
}

// UpdateMovies godoc
// @Summary Update a movie
// @Description Update movie by movie id (admin only)
// @Tags Admin
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Param movie body models.UpdatedMovie true "Movie data to update"
// @Success 200 {object} utils.Response
// @Security BearerAuth
// @Router /admin/{id} [patch]
func UpdateMovies(ctx *gin.Context) {
	role, _ := ctx.Get("role")
	if role != "admin" {
		ctx.JSON(http.StatusForbidden, utils.Response{
			Success: false,
			Message: "Forbidden: Only admin can add movies",
		})
		return
	}

	req := models.UpdatedMovie{}
	ctx.ShouldBindJSON(&req)

	id := ctx.Param("id")
	idInt, _ := strconv.Atoi(id)

	if id == "" {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "User ID is required",
		})
		return
	}

	err := models.GetUpdatedMovie(idInt, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to update movie",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Movie Updated Successfuly!",
	})

}

func TicketSales(ctx *gin.Context) {
	role, exist := ctx.Get("role")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Unauthorized!",
		})
		return
	}

	if role != "admin" {
		ctx.JSON(http.StatusForbidden, utils.Response{
			Success: false,
			Message: "Forbidden: Only admin can access this resource",
		})
		return
	}

	key := ctx.Query("filter")

	data, err := models.GetTicketSales(key)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Internal Server Error",	
			Errors: "lorem ipsum",
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Ticket sales fetched successfully",
		Results:    data,
	})
}

