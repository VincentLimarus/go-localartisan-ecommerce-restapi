package responsesDTO

import (
	"time"

	"github.com/google/uuid"
)

type ArtisansResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	ShopName    string    `json:"shop_name"`
	ShopAddress string    `json:"shop_address"`
	Description string    `json:"description"`
	ShopBanner  string    `json:"shop_banner"`
	Rating      float64   `json:"rating"`
	IsActive    bool      `json:"is_active"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`

	// Start of FK
	Products []ProductResponseDTO `json:"products"`
	// End of FK
}