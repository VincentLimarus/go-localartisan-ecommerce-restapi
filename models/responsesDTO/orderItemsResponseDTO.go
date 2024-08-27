package responsesDTO

import (
	"time"

	"github.com/google/uuid"
)

type OrderItemsResponseDTO struct {
	ID           uuid.UUID    `json:"id"`
	ProductID    uuid.UUID    `json:"product_id"`
	OrderID      uuid.UUID    `json:"order_id"`
	Quantity     int       `json:"quantity"`
	PriceAtOrder float64   `json:"price_at_order"`
	IsActive     bool      `json:"is_active"`
	CreatedBy    string    `json:"created_by"`
	UpdatedBy    string    `json:"updated_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}