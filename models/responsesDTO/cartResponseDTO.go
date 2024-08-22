package responsesDTO

import (
	"time"

	"github.com/google/uuid"
)

type CartResponseDTO struct {
	ID        uuid.UUID    `json:"id"`
	UserID    uuid.UUID    `json:"user_id"`
	IsActive  bool      `json:"is_active"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CartInformation []CartInformationResponseDTO `json:"cart_information"`
}
