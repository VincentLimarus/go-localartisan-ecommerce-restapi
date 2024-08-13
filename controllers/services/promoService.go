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

func GetAllPromoService(c *gin.Context)	{
	var GetAllPromoRequestDTO requestsDTO.GetAllPromosRequestDTO
	GetAllPromoRequestDTO.Page, GetAllPromoRequestDTO.Limit, GetAllPromoRequestDTO.OrderBy, GetAllPromoRequestDTO.OrderType = utils.PaginationHandler(GetAllPromoRequestDTO.Page, GetAllPromoRequestDTO.Limit, GetAllPromoRequestDTO.OrderBy, GetAllPromoRequestDTO.OrderType)
	if err := c.ShouldBindWith(&GetAllPromoRequestDTO, binding.Form); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllPromo(GetAllPromoRequestDTO)
	c.JSON(code, output)
}

func GetAllPromoByProductID(c *gin.Context){
	productID := c.Param("id")

	if _, err := uuid.Parse(productID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllPromoByProductID(productID)
	c.JSON(code, output)
}

func GetPromo(c *gin.Context){
	promoID := c.Param("id")

	if _, err := uuid.Parse(promoID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetPromo(promoID)
	c.JSON(code, output)
}

func CreatePromo(c *gin.Context){
	var CreatePromoRequestDTO requestsDTO.CreatePromosRequestDTO

	if err := c.ShouldBindJSON(&CreatePromoRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.CreatePromo(CreatePromoRequestDTO)
	c.JSON(code, output)
}

func UpdatePromo(c *gin.Context){
	var UpdatePromoRequestDTO requestsDTO.UpdatePromosRequestDTO

	if err := c.ShouldBindJSON(&UpdatePromoRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.UpdatePromo(UpdatePromoRequestDTO)
	c.JSON(code, output)
}

func DeletePromo(c *gin.Context){
	var DeletePromoRequestDTO requestsDTO.DeletePromosRequestDTO

	if err := c.ShouldBindJSON(&DeletePromoRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.DeletePromo(DeletePromoRequestDTO)
	c.JSON(code, output)
}

func BasePromoService(router *gin.RouterGroup) {
	router.GET("/promos", GetAllPromoService)
	router.GET("/promos/product/:id", GetAllPromoByProductID)
	router.GET("/promo/:id", GetPromo)	
}

func AuthPromoService(router *gin.RouterGroup) {
	router.POST("/promo/create", CreatePromo)
	router.POST("/promo/update", UpdatePromo)
	router.POST("/promo/delete", DeletePromo)
}
