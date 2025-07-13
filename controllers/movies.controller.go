package controllers

import (
	"be-cinevo/models"
	"be-cinevo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllMovies godoc
// @Summary Get all movies
// @Description Retrieve all movies with optional search and pagination
// @Tags Movies
// @Produce json
// @Param search query string false "Search by movie title"
// @Param page query int false "Page number"
// @Param limit query int false "Limit per page"
// @Param genre query string false "Filter By Genre"
// @Success 200 {object} utils.Response{results=[]models.Movie}
// @Router /movies [get]
func GetAllMovies(ctx *gin.Context) {
	key := ctx.Query("search")
	filter := ctx.Query("genres")
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

	movies, totalMovies, err := models.HandleShowAllMovies(key, limitInt, offset, filter)
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
	movies, err := models.HandleNowShowingMovies()
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
	movies, err := models.HandleUpComingMovies()
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

// controller
// GetMovieDetail godoc
// @Summary Get movie details by ID
// @Description Get detailed information about a movie including genres, directors, and casts
// @Tags Movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} utils.Response{results=dto.MovieResponse}
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /movies/{id} [get]
func GetMovieDetail(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid movie ID",
			Errors:  "ID must be an integer",
		})
		return
	}

	movie, err := models.FindMovieById(id)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "movie not found" {
			status = http.StatusNotFound
		}

		ctx.JSON(status, utils.Response{
			Success: false,
			Message: "Failed to retrieve movie",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Movie retrieved successfully",
		Results: movie,
	})
}
