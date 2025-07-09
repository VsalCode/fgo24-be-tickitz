package middlewares

import (
	"be-cinevo/utils"
	"errors"
	"fmt"
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
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Authorization header missing!",
			})
			ctx.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, "Bearer ")
		if len(tokenParts) != 2 {
			ctx.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Invalid token format!",
			})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimSpace(tokenParts[1])
		blacklisted, err := utils.RedisClient.Exists(ctx, "blacklist:"+tokenString).Result()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.Response{
				Success: false,
				Message: "Internal server error",
			})
			ctx.Abort()
			return
		}

		if blacklisted > 0 {
			ctx.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Token has been revoked!",
			})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			if errors.Is(err, jwt.ErrTokenExpired) {
				ctx.JSON(http.StatusUnauthorized, utils.Response{
					Success: false,
					Message: "Token expired!",
				})
			} else {
				ctx.JSON(http.StatusUnauthorized, utils.Response{
					Success: false,
					Message: "Invalid token!",
				})
			}
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Invalid token claims!",
			})
			ctx.Abort()
			return
		}

		exp, ok := claims["exp"].(float64)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, utils.Response{
				Success: false,
				Message: "Missing token expiration!",
			})
			ctx.Abort()
			return
		}

		ctx.Set("userId", int(claims["userId"].(float64)))
		ctx.Set("role", claims["role"].(string))
		ctx.Set("token", tokenString)
		ctx.Set("exp", exp)

		ctx.Next()
	}
}
