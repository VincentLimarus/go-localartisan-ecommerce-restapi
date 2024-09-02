package helpers

import (
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/models/outputs"
	"localArtisans/models/repositories"
	"localArtisans/models/requestsDTO"
	"localArtisans/models/responsesDTO"

	"github.com/google/uuid"
)

func GetAllOrderItems(GetAllOrderItems requestsDTO.GetAllOrderItemsRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var orderItems []database.OrderItems
	
	if GetAllOrderItems.Limit > 100 {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Limit can't more than 100",
		}
		return 400, output
	}

	offset := (GetAllOrderItems.Page - 1) * GetAllOrderItems.Limit
	order := GetAllOrderItems.OrderBy + " " + GetAllOrderItems.OrderType
	err := db.Offset(offset).Limit(GetAllOrderItems.Limit).Order(order).Find(&orderItems).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if len(orderItems) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Order Items not exist",
		}
		return 404, output
	}

	var totalData int64
	var totalPage int
	db.Model(&database.OrderItems{}).Count(&totalData)
	if totalData%int64(GetAllOrderItems.Limit) == 0 {
		totalPage = int(totalData / int64(GetAllOrderItems.Limit))
	} else {
		totalPage = int(totalData / int64(GetAllOrderItems.Limit)) + 1
	}

	output := outputs.GetAllOrderItemsOutput{}
	output.Page = GetAllOrderItems.Page
	output.Limit = GetAllOrderItems.Limit
	output.OrderBy = GetAllOrderItems.OrderBy
	output.OrderType = GetAllOrderItems.OrderType
	output.Code = 200
	output.Message = "Success"
	output.TotalData = int(totalData)
	output.TotalTake = len(orderItems)
	output.TotalPage = totalPage

	for _, orderItem := range orderItems {
		output.Data = append(output.Data, responsesDTO.OrderItemsResponseDTO{
			ID:           orderItem.ID,
			ProductID:    orderItem.ProductID,	
			OrderID:      orderItem.OrderID,
			Quantity:     orderItem.Quantity,
			PriceAtOrder: orderItem.PriceAtOrder,
			IsActive:     orderItem.IsActive,
			CreatedBy:    orderItem.CreatedBy,
			UpdatedBy:    orderItem.UpdatedBy,
			CreatedAt:    orderItem.CreatedAt,
			UpdatedAt:    orderItem.UpdatedAt,
		})
	}

	return 200, output
}

func GetOrderItemByID(ID string) (int, interface{}){
	var orderItems database.OrderItems
	orderItems, err := repositories.GetAllOrderItemsByOrderItemsID(ID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if orderItems.ID == uuid.Nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Order Item not exist",
		}
		return 404, output
	}

	output := outputs.GetOrderItemByIDOutput{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responsesDTO.OrderItemsResponseDTO{
		ID:           orderItems.ID,
		ProductID:    orderItems.ProductID,
		OrderID:      orderItems.OrderID,
		Quantity:     orderItems.Quantity,
		PriceAtOrder: orderItems.PriceAtOrder,
		IsActive:     orderItems.IsActive,
		CreatedBy:    orderItems.CreatedBy,
		UpdatedBy:    orderItems.UpdatedBy,
		CreatedAt:    orderItems.CreatedAt,
		UpdatedAt:    orderItems.UpdatedAt,
	}
	return 200, output		
}

func GetAllOrderItemsByOrderIDRequestDTO(orderID string) (int, interface{}){
	var orderItems []database.OrderItems
	orderItems, err := repositories.GetAllOrderItemsByOrderIDFromDatabase(orderID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if len(orderItems) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Order Items not exist",
		}
		return 404, output
	}

	output := outputs.GetAllOrderItemsByOrderIDOutput{}
	output.Code = 200
	output.Message = "Success"

	for _, orderItem := range orderItems {
		output.Data = append(output.Data, responsesDTO.OrderItemsResponseDTO{
			ID:           orderItem.ID,
			ProductID:    orderItem.ProductID,
			OrderID:      orderItem.OrderID,
			Quantity:     orderItem.Quantity,
			PriceAtOrder: orderItem.PriceAtOrder,
			IsActive:     orderItem.IsActive,
			CreatedBy:    orderItem.CreatedBy,
			UpdatedBy:    orderItem.UpdatedBy,
			CreatedAt:    orderItem.CreatedAt,
			UpdatedAt:    orderItem.UpdatedAt,
		})
	}
	return 200, output
}

func GetAllOrderItemsByProductID(productID string) (int, interface{}){
	var orderItems []database.OrderItems
	orderItems, err := repositories.GetAllOrderItemsByProductID(productID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if len(orderItems) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Order Items not exist",
		}
		return 404, output
	}

	output := outputs.GetAllOrderItemsByProductIDOutput{}
	output.Code = 200
	output.Message = "Success"

	for _, orderItem := range orderItems {
		output.Data = append(output.Data, responsesDTO.OrderItemsResponseDTO{
			ID:           orderItem.ID,
			ProductID:    orderItem.ProductID,
			OrderID:      orderItem.OrderID,
			Quantity:     orderItem.Quantity,
			PriceAtOrder: orderItem.PriceAtOrder,
			IsActive:     orderItem.IsActive,
			CreatedBy:    orderItem.CreatedBy,
			UpdatedBy:    orderItem.UpdatedBy,
			CreatedAt:    orderItem.CreatedAt,
			UpdatedAt:    orderItem.UpdatedAt,
		})
	}
	return 200, output
}