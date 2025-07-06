package middlewares

import (
	"be-cinevo/utils"
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
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

		_, err := utils.RedisClient.Get(context.Background(), "blacklist:"+tokenString).Result()
		if err != redis.Nil {
			ctx.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Token has been invalidated!",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

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

		claims, ok := rawToken.Claims.(jwt.MapClaims)
		if !ok || !rawToken.Valid {
			ctx.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Token Invalid!",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userIdFloat, exists := claims["userId"]
		if !exists {
			ctx.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Invalid token claims!",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userId := int(userIdFloat.(float64))
		role, exists := claims["role"].(string)
		if !exists {
			ctx.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Invalid token claims!",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("userId", userId)
		ctx.Set("role", role)
		ctx.Next()
	}
}
