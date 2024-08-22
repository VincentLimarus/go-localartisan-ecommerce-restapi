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

func GetAllCartInformation(c *gin.Context) {
	var GetAllCartInformationsRequestDTO requestsDTO.GetAllCartInformationsRequestDTO
	GetAllCartInformationsRequestDTO.Page, GetAllCartInformationsRequestDTO.Limit, GetAllCartInformationsRequestDTO.OrderBy, GetAllCartInformationsRequestDTO.OrderType = utils.PaginationHandler(GetAllCartInformationsRequestDTO.Page, GetAllCartInformationsRequestDTO.Limit, GetAllCartInformationsRequestDTO.OrderBy, GetAllCartInformationsRequestDTO.OrderType)
	if err := c.ShouldBindWith(&GetAllCartInformationsRequestDTO, binding.Form); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllCartInformations(GetAllCartInformationsRequestDTO)
	c.JSON(code, output)
}

func GetAllCartInformationByCartID(c *gin.Context) {
	cartID := c.Param("id")

	if _, err := uuid.Parse(cartID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllCartInformationByCartID(cartID)
	c.JSON(code, output)
}

func AddItemToCart(c *gin.Context) {
	var AddItemToCartRequestDTO requestsDTO.AddItemToCartRequestDTO
	if err := c.ShouldBindJSON(&AddItemToCartRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.AddItemToCart(AddItemToCartRequestDTO)
	c.JSON(code, output)
}

func UpdateItemInCart(c *gin.Context) {
	var UpdateItemInCartRequestDTO requestsDTO.UpdateItemInCartRequestDTO
	if err := c.ShouldBindJSON(&UpdateItemInCartRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.UpdateItemInCart(UpdateItemInCartRequestDTO)
	c.JSON(code, output)
}

func DeleteItemInCart(c *gin.Context) {
	var DeleteItemInCartRequestDTO requestsDTO.DeleteItemInCartRequestDTO
	if err := c.ShouldBindJSON(&DeleteItemInCartRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.DeleteItemInCart(DeleteItemInCartRequestDTO)
	c.JSON(code, output)
}

func AuthCartInformationService(router *gin.RouterGroup) {
	router.GET("/carts-information", GetAllCartInformation)
	router.GET("/carts-information/cart/:id", GetAllCartInformationByCartID)
	router.POST("/cart-information/add-item", AddItemToCart)
	router.POST("/cart-information/update-item", UpdateItemInCart)
	router.POST("/cart-information/delete-item", DeleteItemInCart)
}