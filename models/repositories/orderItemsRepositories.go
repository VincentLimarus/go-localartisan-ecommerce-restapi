package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/models/responsesDTO"
)

func GetAllOrderItemsByOrderID(orderID string) ([]responsesDTO.OrderItemsResponseDTO, error) {
	var orderItems []responsesDTO.OrderItemsResponseDTO

	db := configs.GetDB()
	err := db.Table("order_items").Where("order_id = ?", orderID).Find(&orderItems).Error

	if err != nil {
		return orderItems, err
	}
	return orderItems, nil
}

func GetAllOrderItemsByOrderIDFromDatabase(orderID string) ([]database.OrderItems, error) {
	var orderItems []database.OrderItems

	db := configs.GetDB()
	err := db.Table("order_items").Where("order_id = ?", orderID).Find(&orderItems).Error

	if err != nil {
		return orderItems, err
	}
	return orderItems, nil
}

func GetAllOrderItemsByOrderItemsID(orderItemsID string) (database.OrderItems, error) {
	var orderItems database.OrderItems

	db := configs.GetDB()
	err := db.Table("order_items").Where("id = ?", orderItemsID).Find(&orderItems).Error

	if err != nil {
		return orderItems, err
	}
	return orderItems, nil
}

func GetAllOrderItemsByProductID(productID string) ([]database.OrderItems, error) {
	var orderItems []database.OrderItems

	db := configs.GetDB()
	err := db.Table("order_items").Where("product_id = ?", productID).Find(&orderItems).Error

	if err != nil {
		return orderItems, err
	}
	return orderItems, nil
}