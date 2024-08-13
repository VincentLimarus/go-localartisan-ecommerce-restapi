package responsesDTO

import (
	"time"

	"github.com/google/uuid"
)

type ProductResponseDTO struct {
	ID           uuid.UUID `json:"id"`
	Name         string   `json:"name"`
	Price        float64  `json:"price"`
	Description  string   `json:"description"`
	Quantity     int      `json:"quantity"`
	ItemSold     int      `json:"item_sold"`
	Rating       float64  `json:"rating"`
	IsActive     bool     `json:"is_active"`
	CreatedBy    string   `json:"created_by"`
	UpdatedBy    string   `json:"updated_by"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	CategoryID   uuid.UUID   `json:"category_id"`
	ArtisanID    uuid.UUID   `json:"artisan_id"` 
}