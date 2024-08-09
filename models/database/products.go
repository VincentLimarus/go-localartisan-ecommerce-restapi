package database

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name        string     `gorm:"type:varchar(255);not null"`
	Price       float64    `gorm:"type:decimal(10,2);not null"`
	Description string     `gorm:"type:text;not null"`
	Quantity    int        `gorm:"type:int;not null"`
	ItemSold    int        `gorm:"type:int;not null;default:0"`
	Rating      float64    `gorm:"type:decimal(2,1);not null;default:0.0"`
	IsActive    bool       `gorm:"type:boolean;not null;default:true"`
	CreatedBy   string     `gorm:"type:varchar(255);not null; default:'system'"`
	UpdatedBy   string     `gorm:"type:varchar(255);not null; default:'system'"`
	CreatedAt   time.Time  `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt   time.Time  `gorm:"type:timestamp;not null;default:now()"`

	// Start of FK
	CategoryID  uuid.UUID  `gorm:"type:uuid;not null"`
	Category    Categories   `gorm:"foreignKey:CategoryID"`

	ArtisanID   uuid.UUID  `gorm:"type:uuid;not null"`
	Artisans     Artisans    `gorm:"foreignKey:ArtisanID"`
	// End of FK
}
