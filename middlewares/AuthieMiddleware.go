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
			log.Println("Error retrieving token:", err)
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				outputs.UnauthorizedOutput{
					Code:    401,
					Message: "Unauthorized token not found",
				},
			)
			return
		}

		log.Println("Received Token:", tokenString)

		is_valid, email, err := utils.ValidateJWTToken(tokenString)
		if err != nil || !is_valid {
			log.Println("Token validation error:", err)
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				outputs.UnauthorizedOutput{
					Code:    401,
					Message: "Unauthorized token is expired or invalid",
				},
			)
			return
		}

		log.Println("Valid Token for Email:", email)

		// Simpan email pengguna dalam context
		c.Set("user_email", email)
		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}
