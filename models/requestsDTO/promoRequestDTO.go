package requestsDTO

import "github.com/google/uuid"

type GetAllPromosRequestDTO struct {
	Page      int    `json:"page" form:"page" binding:"omitempty"`
	Limit     int    `json:"limit" form:"limit" binding:"omitempty"`
	OrderBy   string `json:"order_by" form:"order_by" binding:"omitempty"`
	OrderType string `json:"order_type" form:"order_type" binding:"omitempty"`
}

type GetPromosRequestDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type GetAllPromosByProductIDRequestDTO struct {
	ProductID string `json:"product_id" form:"product_id" binding:"required"`
}

type CreatePromosRequestDTO struct {
	ProductID     uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
	Name          string    `json:"name" form:"name" binding:"required"`
	Description   string    `json:"description" form:"description" binding:"required"`
	PromoDiscount float64   `json:"promo_discount" form:"promo_discount" binding:"required"`
	CreatedBy     string    `json:"created_by" form:"created_by" binding:"omitempty"`
	IsActive      bool      `json:"is_active" form:"is_active" binding:"omitempty"`
}

type UpdatePromosRequestDTO struct {
	ID            string  `json:"id" form:"id" binding:"required"`
	ProductID     uuid.UUID  `json:"product_id" form:"product_id" binding:"omitempty"`
	Name          string  `json:"name" form:"name" binding:"omitempty"`
	Description   string  `json:"description" form:"description" binding:"omitempty"`
	PromoDiscount float64 `json:"promo_discount" form:"promo_discount" binding:"omitempty"`
	UpdatedBy     string  `json:"updated_by" form:"updated_by" binding:"omitempty"`
	IsActive      bool    `json:"is_active" form:"is_active" binding:"omitempty"`
}

type DeletePromosRequestDTO struct {
	// Only admin can delete (Artisan).
	ID string `json:"id" form:"id" binding:"required"`
}