package controllers

import (
	"be-cinevo/models"
	"be-cinevo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllMovies(ctx *gin.Context) {
	key := ctx.Query("search")
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "5")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	if pageInt < 1 || limitInt < 1 {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid page or limit number!",
		})
		return
	}

	offset := (pageInt - 1) * limitInt

	movies, totalMovies, err := models.FindMovieByName(key, limitInt, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to retrieve movies",
			Errors:  err.Error(),
		})
		return
	}

	totalPages := 1
	if limitInt > 0 {
		totalPages = (totalMovies + limitInt - 1) / limitInt
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Movies retrieved successfully",
		Results: movies,
		PageInfo: map[string]interface{}{
			"total":      totalMovies,
			"page":       pageInt,
			"limit":      limitInt,
			"totalPages": totalPages,
		},
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
func GetNowShowingMovies(ctx *gin.Context) {
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
func GetUpComingMovies(ctx *gin.Context) {
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
