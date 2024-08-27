package requestsDTO

import "github.com/google/uuid"

type GetAllCartsRequestDTO struct {
	Page      int    `json:"page" form:"page" binding:"omitempty"`
	Limit     int    `json:"limit" form:"limit" binding:"omitempty"`
	OrderBy   string `json:"order_by" form:"order_by" binding:"omitempty"`
	OrderType string `json:"order_type" form:"order_type" binding:"omitempty"`
}

type GetAllCartByUserIDRequestDTO struct {
	UserID uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
}

type GetCartRequestDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type CreateCartRequestDTO struct {
	UserID    uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	CreatedBy string `json:"created_by" form:"created_by" binding:"omitempty"`
	IsActive  bool   `json:"is_active" form:"is_active" binding:"required"`
}

type DeleteCartRequestDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type OrderNowRequestDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
}
