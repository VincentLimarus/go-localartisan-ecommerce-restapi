package requestsDTO

import "github.com/google/uuid"

type GetAllReviewsRequestDTO struct {
	Page      int    `json:"page" form:"page" binding:"omitempty"`
	Limit     int    `json:"limit" form:"limit" binding:"omitempty"`
	OrderBy   string `json:"order_by" form:"order_by" binding:"omitempty"`
	OrderType string `json:"order_type" form:"order_type" binding:"omitempty"`
}

type GetAllReviewsByProductIDRequestDTO struct {
	ProductID uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
	Page      int    	`json:"page" form:"page" binding:"omitempty"`
	Limit     int    	`json:"limit" form:"limit" binding:"omitempty"`
	OrderBy   string 	`json:"order_by" form:"order_by" binding:"omitempty"`
	OrderType string 	`json:"order_type" form:"order_type" binding:"omitempty"`
}

type GetReviewRequestDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type CreateReviewRequestDTO struct {
	ProductID uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
	UserID    uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	Comment   string    `json:"comment" form:"comment" binding:"required"`
	Rating    float64   `json:"rating" form:"rating" binding:"required"`
	IsActive  bool      `json:"is_active" form:"is_active" binding:"omitempty"`
	CreatedBy string    `json:"created_by" form:"created_by" binding:"omitempty"`
}

type DeleteReviewRequestDTO struct {
	ID              string `json:"id" form:"id" binding:"required"`
	UserID 			uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
}