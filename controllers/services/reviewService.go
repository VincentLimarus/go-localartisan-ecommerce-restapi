package services

import (
	"fmt"
	"localArtisans/controllers/helpers"
	"localArtisans/models/outputs"
	"localArtisans/models/requestsDTO"
	"localArtisans/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

func GetAllReviews(c *gin.Context) {
	var GetAllReviewsRequestDTO requestsDTO.GetAllReviewsRequestDTO
	GetAllReviewsRequestDTO.Page, GetAllReviewsRequestDTO.Limit, GetAllReviewsRequestDTO.OrderBy, GetAllReviewsRequestDTO.OrderType = utils.PaginationHandler(GetAllReviewsRequestDTO.Page, GetAllReviewsRequestDTO.Limit, GetAllReviewsRequestDTO.OrderBy, GetAllReviewsRequestDTO.OrderType)
	if err := c.ShouldBindWith(&GetAllReviewsRequestDTO, binding.Form); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllReviews(GetAllReviewsRequestDTO)
	c.JSON(code, output)
}

func GetAllReviewsByProductID(c *gin.Context) {
	productID := c.Param("id")

	if _, err := uuid.Parse(productID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllReviewsByProductID(productID)
	c.JSON(code, output)
}

func CreateReview(c *gin.Context){
	var CreateReviewRequestDTO requestsDTO.CreateReviewRequestDTO

	if err := c.ShouldBindJSON(&CreateReviewRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.CreateReview(CreateReviewRequestDTO)
	c.JSON(code, output)
}

func DeleteReview(c *gin.Context){
	var DeleteReviewRequestDTO requestsDTO.DeleteReviewRequestDTO

	if err := c.ShouldBindJSON(&DeleteReviewRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.DeleteReview(DeleteReviewRequestDTO)
	c.JSON(code, output)
}

func GetReviewByID(c *gin.Context){
	reviewID := c.Param("id")

	if _, err := uuid.Parse(reviewID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetReviewByID(reviewID)
	c.JSON(code, output)
}


func BaseReviewService(router *gin.RouterGroup) {
	router.GET("/reviews", GetAllReviews)
	router.GET("/reviews/product/:id", GetAllReviewsByProductID)
	router.GET("/review/:id", GetReviewByID)
}

func AuthReviewService(router *gin.RouterGroup) {
	router.POST("/review/create", CreateReview)
	router.POST("/review/delete", DeleteReview)
}
