package responsesDTO

import (
	"time"

	"github.com/google/uuid"
)

type OrderResponseDTO struct {
	ID              uuid.UUID				`json:"id"`
	UserID          uuid.UUID               `json:"user_id"`
	Status          string                  `json:"status"`
	TotalPrice      float64                 `json:"total_price"`
	ShippingAddress string                  `json:"shipping_address"`
	PaymentMethod   string                  `json:"payment_method"`
	IsActive        bool                    `json:"is_active"`
	CreatedBy       string                  `json:"created_by"`
	UpdatedBy       string                  `json:"updated_by"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`

	OrderItems      []OrderItemsResponseDTO `json:"order_items"`
}