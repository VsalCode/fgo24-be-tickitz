package controllers

import (
	"be-cinevo/dto"
	"be-cinevo/models"
	"be-cinevo/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags Auth
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

	err = models.GetNewUser(req)

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
	})

}

// @Summary Login a user
// @Description Login User
// @Tags Auth
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
		return
	}

	token, err := utils.GenerateToken(userId, role)
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

// @Summary Forgot Password
// @Description Get OTP
// @Tags Auth
// @Accept json
// @Produce json
// @Param email body dto.VerifyEmail true "Email address to send verification code"
// @Success 200 {object} utils.Response
// @Router /auth/forgot-password [post]
func ForgotPassword(ctx *gin.Context) {
	req := dto.VerifyEmail{}
	ctx.ShouldBindJSON(&req)

	result, err := models.SendVerificationCode(req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to send verification code",
			Errors:  err.Error(),
		})
		return
	}

	err = utils.SendEmailOTP(req.Email, result.VerificationCode)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to send email",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Verification code sended to your email!",
	})

}

// @Summary Reset Password
// @Description Reset user password using verification code
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.ForgotPasswordRequest true "Forgot password request"
// @Success 200 {object} utils.Response
// @Router /auth/reset-password [post]
func ResetPassword(ctx *gin.Context) {
	req := dto.ForgotPasswordRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid request format",
			Errors:  err.Error(),
		})
		return
	}

	err = models.SendNewPassword(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Failed to reset password",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Password reset successful",
	})
}

func LogoutUser(ctx *gin.Context) {
	token, exists := ctx.Get("token")
	if !exists {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Token not found in context",
		})
		return
	}
	
	tokenString := token.(string)
	
	exp, exists := ctx.Get("exp")
	if !exists {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Expiration not found in context",
		})
		return
	}
	
	expTime := time.Unix(int64(exp.(float64)), 0)
	now := time.Now()
	
	ttl := expTime.Sub(now)
	if ttl < 0 {
		ttl = 1 * time.Minute 
	}
	
	err := utils.RedisClient.Set(ctx, "blacklist:"+tokenString, "revoked", ttl).Err()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "Failed to revoke token",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Logout successful",
	})
}