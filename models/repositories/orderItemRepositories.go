package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/database"
)

func GetOrderItemByOrderID(orderID string) (database.OrderItems, error) {
	var orderItems database.OrderItems
	db := configs.GetDB()

	err := db.Table("order_items").Where("order_id = ?", orderID).First(&orderItems).Error

	if err != nil {
		return orderItems, err
	}
	return orderItems, nil
}