package responsesDTO

import (
	"time"

	"github.com/google/uuid"
)

type CartInformationResponseDTO struct {
	CartID       uuid.UUID    `json:"cartID"`
	ProductID    uuid.UUID    `json:"productID"`
	Quantity     int       `json:"quantity"`
	PriceAtOrder float64   `json:"priceAtOrder"`
	IsActive     bool      `json:"isActive"`
	CreatedBy    string    `json:"createdBy"`
	UpdatedBy    string    `json:"updatedBy"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
