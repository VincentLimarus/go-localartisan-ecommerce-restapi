package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/database"
)

func GetArtisanByArtisanID(artisanID string) (database.Artisans, error){
	var artisan database.Artisans
	
	db := configs.GetDB()
	err := db.Table("artisans").Where("id = ?", artisanID).First(&artisan).Error

	if err != nil {
		return artisan, err
	}
	return artisan, nil
}


func GetAllArtisanByUserID(userID string) ([]database.Artisans, error){
	var artisan []database.Artisans
	
	db := configs.GetDB()
	err := db.Table("artisans").Where("user_id = ?", userID).Find(&artisan).Error

	if err != nil {
		return artisan, err
	}
	return artisan, nil
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