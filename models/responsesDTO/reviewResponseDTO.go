package responsesDTO

import (
	"time"

	"github.com/google/uuid"
)

type ReviewsResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	UserID    uuid.UUID `json:"user_id"`
	Comment   string    `json:"comment"`
	Rating    float64   `json:"rating"`
	IsActive  bool      `json:"is_active"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}