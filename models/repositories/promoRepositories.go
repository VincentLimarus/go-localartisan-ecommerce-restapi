package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/models/responsesDTO"
)

func GetPromoByID(promoID string) (database.Promos, error) {
	var promo database.Promos
	db := configs.GetDB()
	err := db.Table("promos").Where("id = ?", promoID).First(&promo).Error

	if err != nil {
		return promo, err
	}
	return promo, nil
}

func GetAllPromoByProductID(productID string) ([]responsesDTO.PromoResponseDTO, error) {
	var promos []responsesDTO.PromoResponseDTO
	db := configs.GetDB()
	err := db.Table("promos").Where("product_id = ?", productID).Find(&promos).Error

	if err != nil {
		return promos, err
	}
	return promos, nil
}

func GetAllPromos() ([]responsesDTO.PromoResponseDTO, error) {
	var promos []responsesDTO.PromoResponseDTO
	db := configs.GetDB()
	err := db.Table("promos").Find(&promos).Error

	if err != nil {
		return promos, err
	}
	return promos, nil
}