package middlewares

import (
	"localArtisans/models/outputs"
	"localArtisans/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequiredAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				outputs.UnauthorizedOutput{
					Code:    401,
					Message: "Unauthorized token not found",
				},
			)
			return
		}
		is_valid, err := utils.ValidateJWTToken(tokenString)
		if err != nil || !is_valid {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				outputs.UnauthorizedOutput{
					Code:    401,
					Message: "Unauthorized is expired or invalid",
				},
			)
			return
		}
		c.Next()
	}
}
