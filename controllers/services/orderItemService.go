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

func GetAllOrderItems(c *gin.Context) {
	var GetAllOrderItemsRequestDTO requestsDTO.GetAllOrderItemsRequestDTO
	GetAllOrderItemsRequestDTO.Page, GetAllOrderItemsRequestDTO.Limit, GetAllOrderItemsRequestDTO.OrderBy, GetAllOrderItemsRequestDTO.OrderType = utils.PaginationHandler(GetAllOrderItemsRequestDTO.Page, GetAllOrderItemsRequestDTO.Limit, GetAllOrderItemsRequestDTO.OrderBy, GetAllOrderItemsRequestDTO.OrderType)
	if err := c.ShouldBindWith(&GetAllOrderItemsRequestDTO, binding.Form); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllOrderItems(GetAllOrderItemsRequestDTO)
	c.JSON(code, output)
}

func GetOrderItemByID(c *gin.Context) {
	orderItemID := c.Param("id")

	if _, err := uuid.Parse(orderItemID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetOrderItemByID(orderItemID)
	c.JSON(code, output)
}

func GetAllOrderItemsByOrderIDRequestDTO(c *gin.Context) {
	orderID := c.Param("id")

	if _, err := uuid.Parse(orderID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllOrderItemsByOrderIDRequestDTO(orderID)
	c.JSON(code, output)
}

func GetAllOrderItemsByProductID(c *gin.Context) {
	productID := c.Param("id")

	if _, err := uuid.Parse(productID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllOrderItemsByProductID(productID)
	c.JSON(code, output)
}

func AuthOrderItemsService(router *gin.RouterGroup) {
	router.GET("/order-items", GetAllOrderItems)
	router.GET("/order-items/:id", GetOrderItemByID)
	router.GET("/order-items/order/:id", GetAllOrderItemsByOrderIDRequestDTO)
	router.GET("/order-items/product/:id", GetAllOrderItemsByProductID)
}