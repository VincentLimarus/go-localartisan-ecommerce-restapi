package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/models/outputs"
	"localArtisans/models/repositories"
	"localArtisans/models/requestsDTO"
	"localArtisans/models/responsesDTO"
	workers "localArtisans/worker"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllOrders(GetAllOrderRequestDTO requestsDTO.GetAllOrderRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var orders []database.Orders

	if GetAllOrderRequestDTO.Limit > 100 {
		output := outputs.BadRequestOutput{
			Code : 400,
			Message : "Bad Request: Limit can't more than 100",
		}
		return 400, output
	}

	offset := (GetAllOrderRequestDTO.Page - 1) * GetAllOrderRequestDTO.Limit
	order := fmt.Sprintf("%s %s", GetAllOrderRequestDTO.OrderBy, GetAllOrderRequestDTO.OrderType)
	err := db.Offset(offset).Limit(GetAllOrderRequestDTO.Limit).Order(order).Find(&orders).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	if len(orders) == 0 {
		output := outputs.NotFoundOutput{
			Code : 404,
			Message : "Not Found",
		}
		return 404, output
	}

	var totalData int64
	var totalPage int 
	db.Model(&database.Orders{}).Count(&totalData)

	if int(totalData) % GetAllOrderRequestDTO.Limit == 0 {
		totalPage = int(totalData) / GetAllOrderRequestDTO.Limit
	} else {
		totalPage = (int(totalData) / GetAllOrderRequestDTO.Limit) + 1
	}

	output := outputs.GetAllOrderOutput{}
	output.Page = GetAllOrderRequestDTO.Page
	output.Limit = GetAllOrderRequestDTO.Limit
	output.OrderBy = GetAllOrderRequestDTO.OrderBy
	output.OrderType = GetAllOrderRequestDTO.OrderType
	output.Code = 200 
	output.Message = "Success"
	output.TotalData = int(totalData)
	output.TotalTake = len(orders)
	output.TotalPage = totalPage

	for _, order := range orders{
		var orderItems []responsesDTO.OrderItemsResponseDTO
		orderItems, err := repositories.GetAllOrderItemsByOrderID(order.ID.String())

		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code : 500,
				Message : "Internal Server Error",
			}
			return 500, output
		}

		output.Data = append(output.Data, responsesDTO.OrderResponseDTO{
			ID : order.ID,
			UserID : order.UserID,
			Status : order.Status,
			TotalPrice : order.TotalPrice,
			ShippingAddress : order.ShippingAddress,
			PaymentMethod : order.PaymentMethod,
			IsActive : order.IsActive,
			CreatedBy : order.CreatedBy,
			UpdatedBy : order.UpdatedBy,
			CreatedAt : order.CreatedAt,
			UpdatedAt : order.UpdatedAt,
			OrderItems : orderItems,
		})
	}
	return 200, output
}

func GetOrderByOrderID(orderID string) (int, interface{}) {
	var order database.Orders
	order, err := repositories.GetOrderByOrderID(orderID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	if len(order.ID) == 0 {
		output := outputs.NotFoundOutput{
			Code : 404,
			Message : "Not Found",
		}
		return 404, output
	}

	var orderItems []responsesDTO.OrderItemsResponseDTO
	orderItems, err = repositories.GetAllOrderItemsByOrderID(order.ID.String())

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	output := outputs.GetOrderOutputByOrderIDOutput{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responsesDTO.OrderResponseDTO{
		ID : order.ID,
		UserID : order.UserID,
		Status : order.Status,
		TotalPrice : order.TotalPrice,
		ShippingAddress : order.ShippingAddress,
		PaymentMethod : order.PaymentMethod,
		IsActive : order.IsActive,
		CreatedBy : order.CreatedBy,
		UpdatedBy : order.UpdatedBy,
		CreatedAt : order.CreatedAt,
		UpdatedAt : order.UpdatedAt,
		OrderItems : orderItems,
	}
	return 200, output
}	

func GetAllOrderByUserID(userID string) (int, interface{}){
	var orders []database.Orders
	orders, err := repositories.GetAllOrderByUserID(userID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	if len(orders) == 0 {
		output := outputs.NotFoundOutput{
			Code : 404,
			Message : "Not Found",
		}
		return 404, output
	}

	output := outputs.GetAllOrderByUserIDOutput{}
	output.Code = 200
	output.Message = "Success"
	
	for _, order := range orders{
		var orderItems []responsesDTO.OrderItemsResponseDTO
		orderItems, err := repositories.GetAllOrderItemsByOrderID(order.ID.String())

		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code : 500,
				Message : "Internal Server Error",
			}
			return 500, output
		}

		output.Data = append(output.Data, responsesDTO.OrderResponseDTO{
			ID : order.ID,
			UserID : order.UserID,
			Status : order.Status,
			TotalPrice : order.TotalPrice,
			ShippingAddress : order.ShippingAddress,
			PaymentMethod : order.PaymentMethod,
			IsActive : order.IsActive,
			CreatedBy : order.CreatedBy,
			UpdatedBy : order.UpdatedBy,
			CreatedAt : order.CreatedAt,
			UpdatedAt : order.UpdatedAt,
			OrderItems : orderItems,
		})
	}
	return 200, output
}

func DeleteOrder(c *gin.Context, DeleteOrderRequestDTO requestsDTO.DeleteOrderRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	redisClient := configs.GetRedis()
	order := database.Orders{}

	err := db.Table("orders").Where("id = ?", DeleteOrderRequestDTO.ID).First(&order).Error
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error",
		}
		return 500, output
	}

	if order.ID == uuid.Nil {
		output := outputs.NotFoundOutput{
			Code:    404,
			Message: "Not Found",
		}
		return 404, output
	}

	var orderItems []responsesDTO.OrderItemsResponseDTO
	orderItems, err = repositories.GetAllOrderItemsByOrderID(order.ID.String())
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error",
		}
		return 500, output
	}

	email, exists := c.Get("user_email")
	if !exists {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error",
		}
		return 500, output
	}
	
	logData := workers.LogData{
		Action:    fmt.Sprintf("Order %s deleted", order.ID),
		UpdatedBy: email.(string),
		UpdatedAt: time.Now(),
	}

	logDataJSON, _ := json.Marshal(logData)

	ctx := context.Background()
	err = redisClient.RPush(ctx, "log_activity_queue", logDataJSON).Err()
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Failed to enqueue log activity",
		}
		return 500, output
	}

	err = db.Table("orders").Where("id = ?", DeleteOrderRequestDTO.ID).Delete(&order).Error
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code:    500,
			Message: "Internal Server Error",
		}
		return 500, output
	}

	output := outputs.DeleteOrderOutput{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responsesDTO.OrderResponseDTO{
		ID : order.ID,
		UserID : order.UserID,
		Status : order.Status,
		TotalPrice : order.TotalPrice,
		ShippingAddress : order.ShippingAddress,
		PaymentMethod : order.PaymentMethod,
		IsActive : order.IsActive,
		CreatedBy : order.CreatedBy,
		UpdatedBy : order.UpdatedBy,
		CreatedAt : order.CreatedAt,
		UpdatedAt : order.UpdatedAt,	
		OrderItems : orderItems,
	}
	return 200, output
}


func PayOrder(PayOrderRequestDTO requestsDTO.PayOrderRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var order database.Orders

	err := db.Table("orders").Where("id = ?", PayOrderRequestDTO.ID).First(&order).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}	

	if len(order.ID) == 0 {
		output := outputs.NotFoundOutput{
			Code : 404,
			Message : "Not Found",
		}
		return 404, output
	}

	if PayOrderRequestDTO.PaymentMethod != ""{
		order.PaymentMethod = PayOrderRequestDTO.PaymentMethod
	}

	if (PayOrderRequestDTO.ConfirmOrder) {
		order.Status = "Order Paid"
	}

	err = db.Save(&order).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	var items []database.OrderItems
	err = db.Table("order_items").Where("order_id = ?", PayOrderRequestDTO.ID).Find(&items).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	for _, item := range items {
		var product database.Products
		err = db.Table("products").Where("id = ?", item.ProductID).First(&product).Error

		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code : 500,
				Message : "Internal Server Error",
			}
			return 500, output
		}

		product.Quantity = product.Quantity - item.Quantity
		err = db.Save(&product).Error

		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code : 500,
				Message : "Internal Server Error",
			}
			return 500, output
		}
	}

	var orderItems []responsesDTO.OrderItemsResponseDTO
	orderItems, err = repositories.GetAllOrderItemsByOrderID(order.ID.String())

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	output := outputs.PayOrderOutput{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responsesDTO.OrderResponseDTO{
		ID : order.ID,
		UserID : order.UserID,
		Status : order.Status,
		TotalPrice : order.TotalPrice,
		ShippingAddress : order.ShippingAddress,
		PaymentMethod : order.PaymentMethod,
		IsActive : order.IsActive,
		CreatedBy : order.CreatedBy,
		UpdatedBy : order.UpdatedBy,
		CreatedAt : order.CreatedAt,	
		UpdatedAt : order.UpdatedAt,
		OrderItems : orderItems,
	}
	return 200, output
}

func GetAllOrderByUserIDAndStatus(GetAllOrderByUserIDAndStatusRequestDTO requestsDTO.GetAllOrderByUserIDAndStatusRequestDTO) (int, interface{}) {
	var orders []database.Orders
	orders, err := repositories.GetAllOrderByUserIDAndStatus(GetAllOrderByUserIDAndStatusRequestDTO.UserID, GetAllOrderByUserIDAndStatusRequestDTO.Status)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	if len(orders) == 0 {
		output := outputs.NotFoundOutput{
			Code : 404,
			Message : "Not Found",
		}
		return 404, output
	}

	output := outputs.GetAllOrderByUserIDAndStatusOutput{}
	output.Code = 200
	output.Message = "Success"
	
	for _, order := range orders{
		var orderItems []responsesDTO.OrderItemsResponseDTO
		orderItems, err := repositories.GetAllOrderItemsByOrderID(order.ID.String())

		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code : 500,
				Message : "Internal Server Error",
			}
			return 500, output
		}

		output.Data = append(output.Data, responsesDTO.OrderResponseDTO{
			ID : order.ID,
			UserID : order.UserID,
			Status : order.Status,
			TotalPrice : order.TotalPrice,
			ShippingAddress : order.ShippingAddress,
			PaymentMethod : order.PaymentMethod,
			IsActive : order.IsActive,
			CreatedBy : order.CreatedBy,
			UpdatedBy : order.UpdatedBy,
			CreatedAt : order.CreatedAt,
			UpdatedAt : order.UpdatedAt,
			OrderItems : orderItems,
		})
	}
	return 200, output
}

func FinishOrder(FinishOrderRequestDTO requestsDTO.FinishOrderRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var order database.Orders

	err := db.Table("orders").Where("id = ?", FinishOrderRequestDTO.ID).First(&order).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}	

	if len(order.ID) == 0 {
		output := outputs.NotFoundOutput{
			Code : 404,
			Message : "Not Found",
		}
		return 404, output
	}

	order.Status = "Order Finished"
	err = db.Save(&order).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	var orderItems []responsesDTO.OrderItemsResponseDTO
	orderItems, err = repositories.GetAllOrderItemsByOrderID(order.ID.String())

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	output := outputs.FinishOrderOutput{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responsesDTO.OrderResponseDTO{
		ID : order.ID,
		UserID : order.UserID,
		Status : order.Status,
		TotalPrice : order.TotalPrice,
		ShippingAddress : order.ShippingAddress,
		PaymentMethod : order.PaymentMethod,
		IsActive : order.IsActive,
		CreatedBy : order.CreatedBy,
		UpdatedBy : order.UpdatedBy,
		CreatedAt : order.CreatedAt,	
		UpdatedAt : order.UpdatedAt,
		OrderItems : orderItems,
	}
	return 200, output
}

func CancelOrder(CancelOrderRequestDTO requestsDTO.CancelOrderRequestDTO) (int, interface{}){
	db := configs.GetDB()
	var order database.Orders

	err := db.Table("orders").Where("id = ?", CancelOrderRequestDTO.ID).First(&order).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	if len(order.ID) == 0 {
		output := outputs.NotFoundOutput{
			Code : 404,
			Message : "Not Found",
		}
		return 404, output
	}

	order.Status = "Order Cancelled"
	err = db.Save(&order).Error
	
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	// Quantity tambahin lagi ke product
	var items []database.OrderItems
	err = db.Table("order_items").Where("order_id = ?", CancelOrderRequestDTO.ID).Find(&items).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	for _, item := range items {
		var product database.Products
		err = db.Table("products").Where("id = ?", item.ProductID).First(&product).Error
		
		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code : 500,
				Message : "Internal Server Error",
			}
			return 500, output
		}

		product.Quantity = product.Quantity + item.Quantity
		err = db.Save(&product).Error

		if err != nil {
			output := outputs.InternalServerErrorOutput{
				Code : 500,
				Message : "Internal Server Error",
			}
			return 500, output
		}
	}
	
	var orderItems []responsesDTO.OrderItemsResponseDTO
	orderItems, err = repositories.GetAllOrderItemsByOrderID(order.ID.String())

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code : 500,
			Message : "Internal Server Error",
		}
		return 500, output
	}

	output := outputs.CancelOrderOutput{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responsesDTO.OrderResponseDTO{
		ID : order.ID,
		UserID : order.UserID,
		Status : order.Status,
		TotalPrice : order.TotalPrice,
		ShippingAddress : order.ShippingAddress,
		PaymentMethod : order.PaymentMethod,
		IsActive : order.IsActive,
		CreatedBy : order.CreatedBy,
		UpdatedBy : order.UpdatedBy,
		CreatedAt : order.CreatedAt,
		UpdatedAt : order.UpdatedAt,
		OrderItems : orderItems,
	}
	return 200, output
}