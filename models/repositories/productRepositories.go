package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/models/responsesDTO"
)

func GetAllProductByArtisanID(artisanID string) ([]responsesDTO.ProductResponseDTO, error) {
	var product []responsesDTO.ProductResponseDTO

	db := configs.GetDB()
	err := db.Table("products").Where("artisan_id = ?", artisanID).Find(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}


func GetAllProductByCategoryID(categoryID string) ([]responsesDTO.ProductResponseDTO, error) {
	var product []responsesDTO.ProductResponseDTO

	db := configs.GetDB()
	err := db.Table("products").Where("category_id = ?", categoryID).Find(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}

func GetProductByProductID(productID string) (database.Products, error) {
	var product database.Products

	db := configs.GetDB()
	err := db.Table("products").Where("id = ?", productID).First(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}


func GetProductByCategoryID(categoryID string) (database.Products, error) {
	var product database.Products

	db := configs.GetDB()
	err := db.Table("products").Where("category_id = ?", categoryID).First(&product).Error

	if err != nil {
		return product, err
	}
	return product, nil
}