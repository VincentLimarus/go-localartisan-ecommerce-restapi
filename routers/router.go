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

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"Code": 400, "message": "Route not found"})
	})

	router.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Code": 400, "message": "Welcome to Local Artisans API"})
	})

	router.GET("/help", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Code": 400, "message": "This is the local artisans API V1 Documentation, for more information please visit /api/v1/docs"})
	})

	router.GET("/api/v1/docs", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Code": 400, "message": "Documentation will be added soon"}) // Docs will be added later
	})

	// --------------- Base Routes ----------------
	base := router.Group("api/v1")
	services.BaseUserService(base)
	services.BaseProductService(base)
	services.BaseArtisanService(base)
	services.BaseCategoryService(base)
	services.BasePromoService(base)
	services.BaseReviewService(base)

	// --------------- Authenticated Routes ----------------
	auth := router.Group("api/v1")
	auth.Use(middlewares.RequiredAuth())
	services.AuthUserService(auth)
	services.AuthProductService(auth)
	services.AuthArtisanService(auth)
	services.AuthCategoryService(auth)
	services.AuthPromoService(auth)
	services.AuthReviewService(auth)
	services.AuthCartService(auth)

	return router
}
