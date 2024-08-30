package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/models/responsesDTO"
)

func GetAllReviewsByProductID(productID string) ([]responsesDTO.ReviewsResponseDTO, error) {
	var reviews []responsesDTO.ReviewsResponseDTO

	db := configs.GetDB()
	err := db.Table("reviews").Where("product_id = ?", productID).Find(&reviews).Error

	if err != nil {
		return reviews, err
	}
	return reviews, nil
}

func GetAllReviewsByUserID(userID string) ([]responsesDTO.ReviewsResponseDTO, error) {
	var reviews []responsesDTO.ReviewsResponseDTO

	db := configs.GetDB()
	err := db.Table("reviews").Where("user_id = ?", userID).Find(&reviews).Error

	if err != nil {
		return reviews, err
	}
	return reviews, nil
}

func GetReviewByID(reviewID string) (database.Reviews, error) {
	var review database.Reviews

	db := configs.GetDB()
	err := db.Table("reviews").Where("id = ?", reviewID).Find(&review).Error

	if err != nil {
		return review, err
	}
	return review, nil
}