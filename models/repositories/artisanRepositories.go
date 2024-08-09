package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/database"
)

func GetArtisanByArtisanID(artisanID string) (database.Artisans, error){
	var artisan database.Artisans
	
	db := configs.GetDB()
	err := db.Where("artisan_id = ?", artisanID).First(&artisan).Error

	if err != nil {
		return artisan, err
	}
	return artisan, nil
}