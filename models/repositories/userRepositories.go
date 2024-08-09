package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/database"
)

func GetUserByUserID(UserID string) (database.User, error) {
	var user database.User

	db := configs.GetDB()
	err := db.Table("users").Where("id = ?", UserID).First(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func GetArtisanByUserID(UserID string) (database.Artisans, error) {
	var artisan database.Artisans

	db := configs.GetDB()
	err := db.Table("artisans").Where("user_id = ?", UserID).First(&artisan).Error

	if err != nil {
		return artisan, err
	}
	return artisan, nil
}
