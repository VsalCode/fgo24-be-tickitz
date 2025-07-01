package middlewares

import (
	"be-cinevo/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		godotenv.Load()
		secretKey := os.Getenv("APP_SECRET")
		token := strings.Split(ctx.GetHeader("Authorization"), "Bearer ")

		if len(token) < 2 {
			ctx.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Unauthorized!",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimSpace(token[1])
		rawToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil {
			if strings.Contains(err.Error(), "expired") {
				ctx.JSON(http.StatusUnauthorized, utils.Response{
					Success: false,
					Message: "Token Expired!",
				})
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Token Invalid!",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userIdFloat := rawToken.Claims.(jwt.MapClaims)["userId"]
		userId := int(userIdFloat.(float64))

		ctx.Set("userId", userId)
		ctx.Next()
	}
}
