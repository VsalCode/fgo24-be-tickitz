package controllers

import (
	"be-cinevo/models"
	"be-cinevo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllMovies(ctx *gin.Context) {
	movies, err := models.FindAllMovies()
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


func GetNowShowingMovies(ctx *gin.Context){}

func GetUpComingMovies(ctx *gin.Context){}