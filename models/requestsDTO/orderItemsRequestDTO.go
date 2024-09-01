package requestsDTO

import "github.com/google/uuid"

type GetAllOrderItemsRequestDTO struct {
	Page      int    `json:"page" form:"page" binding:"omitempty"`
	Limit     int    `json:"limit" form:"limit" binding:"omitempty"`
	OrderBy   string `json:"order_by" form:"order_by" binding:"omitempty"`
	OrderType string `json:"order_type" form:"order_type" binding:"omitempty"`
}

type GetOrderItemByIDRequestDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type GetAllOrderItemsByOrderIDRequestDTO struct {
	OrderID uuid.UUID `json:"order_id" form:"order_id" binding:"required"`
}

type GetAllOrderItemsByProductIDRequestDTO struct {
	ProductID uuid.UUID `json:"product_id" form:"product_id" binding:"required"`
}
