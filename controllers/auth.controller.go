package controllers

import (
	"be-cinevo/dto"
	"be-cinevo/models"
	"be-cinevo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @Summary Register a new user
// @Description Register a new user with email and password
// @Accept json
// @Produce json
// @Param user body dto.RegisterRequest true "User registration details"
// @Success 201 {object} utils.Response{results=models.User}
// @Router /auth/register [post]
func RegisterUser(ctx *gin.Context) {

	req := dto.RegisterRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request format",
			Errors:  err.Error(),
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

	ctx.JSON(http.StatusCreated, utils.Response{
		Success: true,
		Message: "User registered successfully",
		Results: result,
	})

}

// @Summary Login a user
// @Description Login User
// @Accept json
// @Produce json
// @Param user body dto.LoginRequest true "User login details"
// @Success 200 {object} utils.Response{results=string}
// @Router /auth/login [post]
func LoginUser(ctx *gin.Context) {
	godotenv.Load()

	req := dto.LoginRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request format",
			Errors:  err.Error(),
		})
		return
	}

	userId, role, err := models.ValidateLogin(req)	
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, utils.Response{
			Success: false,
			Message: "Invalid email or password",
			Errors:  err.Error(),
		})
	}

	token, err := generateToken(userId, role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to generate token",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Login successful",
		Results: token,
	})
}
