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

func GetAllOrder(c *gin.Context) {
	var GetAllOrderRequestDTO requestsDTO.GetAllOrderRequestDTO
	GetAllOrderRequestDTO.Page, GetAllOrderRequestDTO.Limit, GetAllOrderRequestDTO.OrderBy, GetAllOrderRequestDTO.OrderType = utils.PaginationHandler(GetAllOrderRequestDTO.Page, GetAllOrderRequestDTO.Limit, GetAllOrderRequestDTO.OrderBy, GetAllOrderRequestDTO.OrderType)
	if err := c.ShouldBindWith(&GetAllOrderRequestDTO, binding.Form); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllOrders(GetAllOrderRequestDTO)
	c.JSON(code, output)
}

func GetOrderByOrderID(c *gin.Context) {
	orderID := c.Param("id")

	if _, err := uuid.Parse(orderID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetOrderByOrderID(orderID)
	c.JSON(code, output)
}

func GetAllOrderByUserID(c *gin.Context) {
	userID := c.Param("id")

	if _, err := uuid.Parse(userID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllOrderByUserID(userID)
	c.JSON(code, output)
}

func GetAllOrderByUserIDAndStatus(c *gin.Context) {
	var GetAllOrderByUserIDAndStatusRequestDTO requestsDTO.GetAllOrderByUserIDAndStatusRequestDTO
	if err := c.ShouldBindJSON(&GetAllOrderByUserIDAndStatusRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetAllOrderByUserIDAndStatus(GetAllOrderByUserIDAndStatusRequestDTO)
	c.JSON(code, output)
}

func DeleteOrder(c *gin.Context) {
	var DeleteOrderRequestDTO requestsDTO.DeleteOrderRequestDTO
	if err := c.ShouldBindJSON(&DeleteOrderRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.DeleteOrder(DeleteOrderRequestDTO)
	c.JSON(code, output)
}

func PayOrder(c *gin.Context) {
	var PayOrderRequestDTO requestsDTO.PayOrderRequestDTO
	if err := c.ShouldBindJSON(&PayOrderRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.PayOrder(PayOrderRequestDTO)
	c.JSON(code, output)
}

func FinishOrder(c *gin.Context) {
	var FinishOrderRequestDTO requestsDTO.FinishOrderRequestDTO
	if err := c.ShouldBindJSON(&FinishOrderRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.FinishOrder(FinishOrderRequestDTO)
	c.JSON(code, output)
}

func CancelOrder(c *gin.Context) {
	var CancelOrderRequestDTO requestsDTO.CancelOrderRequestDTO
	if err := c.ShouldBindJSON(&CancelOrderRequestDTO); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: fmt.Sprintf("Bad Request: %v", err),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.CancelOrder(CancelOrderRequestDTO)
	c.JSON(code, output)
}

func AuthOrderService(router *gin.RouterGroup) {
	router.GET("/orders", GetAllOrder)
	router.GET("/order/:id", GetOrderByOrderID)
	router.GET("/orders/user/:id", GetAllOrderByUserID)
	router.POST("/orders/user/status", GetAllOrderByUserIDAndStatus)
	router.POST("/order/delete", DeleteOrder)
	router.POST("/order/pay", PayOrder)
	router.POST("/order/finish", FinishOrder)
	router.POST("/order/cancel", CancelOrder)
}