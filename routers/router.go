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

	router.GET("/help", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Code": 400, "message": "This is the local artisans API V1 Documentation, for more information please visit /api/v1/docs"}) // Docs will be added later
	})

	base := router.Group("api/v1")
	services.UserService(base)
	
	return router
}
