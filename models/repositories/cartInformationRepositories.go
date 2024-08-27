package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/models/responsesDTO"
)

func GetAllCartInformationsByCartID(cartID string) ([]responsesDTO.CartInformationResponseDTO, error) {
	var cartInformations []responsesDTO.CartInformationResponseDTO

	db := configs.GetDB()
	err := db.Table("cart_informations").Where("cart_id = ?", cartID).Find(&cartInformations).Error

	if err != nil {
		return cartInformations, err
	}
	return cartInformations, nil
}

func GetCartInformationByCartIDAndProductID(cartID string, productID string) (responsesDTO.CartInformationResponseDTO, error) {
	var cartInformation responsesDTO.CartInformationResponseDTO

	db := configs.GetDB()
	err := db.Table("cart_informations").Where("cart_id = ? AND product_id = ?", cartID, productID).First(&cartInformation).Error

	if err != nil {
		return cartInformation, err
	}
	return cartInformation, nil
}

func GetAllCartInformationsByCartIDForCartInformation(cartID string) ([]database.CartInformations, error) {
	var cartInformations []database.CartInformations

	db := configs.GetDB()
	err := db.Table("cart_informations").Where("cart_id = ?", cartID).Find(&cartInformations).Error

	if err != nil {
		return cartInformations, err
	}
	return cartInformations, nil
}

func GetCartByID(cartID string) (database.Carts, error) {
	var cart database.Carts

	db := configs.GetDB()
	err := db.Table("carts").Where("id = ?", cartID).First(&cart).Error

	if err != nil {
		return cart, err
	}
	return cart, nil
}
