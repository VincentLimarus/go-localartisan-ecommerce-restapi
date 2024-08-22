package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/database"
)

func GetAllCartsByUserID(userID string) ([]database.Carts, error) {
	var carts []database.Carts

	db := configs.GetDB()
	err := db.Table("carts").Where("user_id = ?", userID).Find(&carts).Error

	if err != nil {
		return carts, err
	}
	return carts, nil
}