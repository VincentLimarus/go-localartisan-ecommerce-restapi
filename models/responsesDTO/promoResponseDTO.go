package responsesDTO

import (
	"time"

	"github.com/google/uuid"
)

type PromoResponseDTO struct {
	ID            uuid.UUID  `json:"id"`
	ProductID     uuid.UUID  `json:"product_id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	PromoDiscount float64   `json:"promo_discount"`
	IsActive      bool      `json:"is_active"`
	CreatedBy     string    `json:"created_by"`
	UpdatedBy     string    `json:"updated_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
