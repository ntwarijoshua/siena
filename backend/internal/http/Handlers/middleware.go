package Handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ntwarijoshua/siena/internal/services"
	"net/http"
	"os"
	"strings"
)

func (app *App) AuthMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		token := strings.Split(c.GetHeader("AUTHORIZATION"), "Bearer ")[1]
		jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, err error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected Signing Method: %v !", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SIGNING_KEY")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
		if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
			userService := app.ServiceContainer.GetService("userService").(*services.UserService)
			if user, err := userService.GetUserByMail(claims["email"].(string)); err == nil {
				c.Set("user", user)
				c.Next()
			}

		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
			return
		}
	}
}
