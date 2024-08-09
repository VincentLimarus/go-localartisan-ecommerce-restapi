package responsesDTO

import (
	"time"

	"github.com/google/uuid"
)

type CategoryResponseDTO struct {
	ID        uuid.UUID    `json:"id"`
	Name      string    `json:"name"`
	Image     string    `json:"image"`
	IsActive  bool      `json:"isActive"`
	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// Start of FK
	Products []ProductResponseDTO `json:"products"`
	// End of FK
}
