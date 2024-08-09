package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/database"
)

func GetUserByArtisanID(artisanID string) (database.User, error) {
	var user database.User

	db := configs.GetDB()
	err := db.Table("users").Where("artisan_id = ?", artisanID).First(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}