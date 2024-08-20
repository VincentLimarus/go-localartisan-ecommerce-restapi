package repositories

import (
	"localArtisans/configs"
	"localArtisans/models/responsesDTO"
)

func GetAllReviewsByProductID(reviewsID string) ([]responsesDTO.ReviewsResponseDTO, error) {
	var reviews []responsesDTO.ReviewsResponseDTO

	db := configs.GetDB()
	err := db.Table("reviews").Where("product_id = ?", reviewsID).Find(&reviews).Error

	if err != nil {
		return reviews, err
	}
	return reviews, nil
}

func GetAllReviewsByUserID(reviewsID string) ([]responsesDTO.ReviewsResponseDTO, error) {
	var reviews []responsesDTO.ReviewsResponseDTO

	db := configs.GetDB()
	err := db.Table("reviews").Where("user_id = ?", reviewsID).Find(&reviews).Error

	if err != nil {
		return reviews, err
	}
	return reviews, nil
}