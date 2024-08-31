package repositories

import (
	"localArtisans/configs"
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