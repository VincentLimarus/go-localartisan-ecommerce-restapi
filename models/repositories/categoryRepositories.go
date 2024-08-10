package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/database"
)


func GetCategoryByCategoryID(categoryID string) (database.Categories, error) {
	var category database.Categories

	db := configs.GetDB()
	err := db.Table("categories").Where("id = ?", categoryID).First(&category).Error

	if err != nil {
		return category, err
	}
	return category, nil
}