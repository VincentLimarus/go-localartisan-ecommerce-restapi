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

func GetUserByEmail(Email string) (database.User, error) {
	var user database.User

	db := configs.GetDB()
	err := db.Table("users").Where("email = ?", Email).First(&user).Error
	
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByAddress(Address string) (database.User, error) {
	var user database.User

	db := configs.GetDB()
	err := db.Table("users").Where("address = ?", Address).First(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}