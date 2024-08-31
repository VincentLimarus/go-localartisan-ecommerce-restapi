package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/database"
)

func GetOrderByOrderID(orderID string) (database.Orders, error) {
	db := configs.GetDB()
	order := database.Orders{}

	err := db.Table("orders").Where("id = ?", orderID).First(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func GetAllOrderByUserID(userID string) ([]database.Orders, error) {
	db := configs.GetDB()
	orders := []database.Orders{}

	err := db.Table("orders").Where("user_id = ?", userID).Find(&orders).Error

	if err != nil {
		return orders, err
	}

	return orders, nil
}