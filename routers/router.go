package routers

import (
	"localArtisans/controllers/services"
	"localArtisans/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutersConfiguration() *gin.Engine{
	router := gin.New()

	router.Use(middlewares.CORSMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// User
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"Code": 400, "message": "Route not found"})
	})

	router.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Code": 400, "message": "Welcome to Local Artisans API"})
	})

	base := router.Group("api/v1")
	services.UserService(base)

	return router
}
