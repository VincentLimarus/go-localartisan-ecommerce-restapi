package requestsDTO

import "github.com/google/uuid"

type GetAllOrderRequestDTO struct {
	Page      int    `json:"page" form:"page" binding:"omitempty"`
	Limit     int    `json:"limit" form:"limit" binding:"omitempty"`
	OrderBy   string `json:"order_by" form:"order_by" binding:"omitempty"`
	OrderType string `json:"order_type" form:"order_type" binding:"omitempty"`
}

type GetOrderRequestDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type GetAllOrderByUserIDRequestDTO struct {
	UserID uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
}

type CreateOrderRequestDTO struct {
	UserID          uuid.UUID `json:"user_id" form:"user_id" binding:"required"`
	Status          string    `json:"status" form:"status" binding:"required"`
	TotalPrice      float64   `json:"total_price" form:"total_price" binding:"required"`
	ShippingAddress string    `json:"shipping_address" form:"shipping_address" binding:"required"`
	PaymentMethod   string    `json:"payment_method" form:"payment_method" binding:"required"`
	CreatedBy       string    `json:"created_by" form:"created_by" binding:"omitempty"`
	IsActive        bool      `json:"is_active" form:"is_active" binding:"required"`
}

type DeleteOrderRequestDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
}

type PayOrderRequestDTO struct {
	ID string `json:"id" form:"id" binding:"required"`
	PaymentMethod string `json:"payment_method" form:"payment_method" binding:"omitempty"` // kalo mau ganti payment bisa. kalo ga diisi, defaultnya tetep payment method yang lama
	ConfirmOrder bool `json:"confirm_order" form:"confirm_order" binding:"required"`		
}