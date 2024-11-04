package middleware

import (
	"fmt"
	"go-clean-todo/infrastructure/mysql/repository"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET")), nil
		})

		if tokenErr != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			userRepository := repository.NewUserRepository()
			user, err := userRepository.FetchByUserID(uint(claims["user_id"].(float64)))
			if err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.Set("user_id", user.UserID())
			// コントローラ処理へ
			ctx.Next()
			// コントローラの後処理が必要な場合はこれ以降に書く
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
