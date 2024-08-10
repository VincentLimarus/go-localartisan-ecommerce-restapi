package database

import (
	"time"

	"github.com/google/uuid"
)

type Artisans struct {
	ID 			uuid.UUID 	`gorm:"primaryKey;type:uuid;default:uuid_generate_v4(); not null"`
	UserID 		uuid.UUID 	`gorm:"type:uuid;not null"`
	ShopName 	string 		`gorm:"type:varchar(255);not null"`
	ShopAddress string 		`gorm:"type:text;"`
	Description string 		`gorm:"type:text;"`
	ShopBanner 	string 		`gorm:"type:varchar(255);"`
	Rating 		float64 	`gorm:"type:float;default:0"`
	IsActive 	bool 		`gorm:"type:boolean;default:true"`
	CreatedBy 	string 		`gorm:"type:varchar(50);not null; default:'system'"`
	UpdatedBy 	string 		`gorm:"type:varchar(50);not null; default:'system'"`
	CreatedAt   time.Time 	`gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt   time.Time 	`gorm:"autoUpdateTime;not null;default:now()"`

	// Start of References
	Product 	[]Product 	`gorm:"foreignKey:ArtisanID"`
	// End of References
}	