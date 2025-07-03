package controllers

import (
	"be-cinevo/models"
	"be-cinevo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllMovies godoc
// @Summary Get all movies
// @Description Retrieve all movies
// @Tags Movies
// @Produce json
// @Success 200 {object} utils.Response{results=[]models.Movie}
// @Failure 500 {object} utils.Response
// @Router /movies [get]
func GetAllMovies(ctx *gin.Context) {
	movies, err := models.FindAllMovies("all")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to retrieve movies",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Movies retrieved successfully",
		Results: movies,
	})
}

// GetNowShowingMovies godoc
// @Summary Get now showing movies
// @Description Retrieve movies that are now showing
// @Tags Movies
// @Produce json
// @Success 200 {object} utils.Response{results=[]models.Movie}
// @Failure 500 {object} utils.Response
// @Router /movies/now-showing [get]
func GetNowShowingMovies(ctx *gin.Context){
	movies, err := models.FindAllMovies("showing")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to retrieve movies",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Movies retrieved successfully",
		Results: movies,
	})
}

// GetUpComingMovies godoc
// @Summary Get upcoming movies
// @Description Retrieve upcoming movies
// @Tags Movies
// @Produce json
// @Success 200 {object} utils.Response{results=[]models.Movie}
// @Failure 500 {object} utils.Response
// @Router /movies/upcoming [get]
func GetUpComingMovies(ctx *gin.Context){
	movies, err := models.FindAllMovies("upcoming")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to retrieve movies",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Movies retrieved successfully",
		Results: movies,
	})
}