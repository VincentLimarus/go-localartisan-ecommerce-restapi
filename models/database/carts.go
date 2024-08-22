package database

import (
	"time"

	"github.com/google/uuid"
)

type Carts struct {
	ID 			uuid.UUID 	`gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID 		uuid.UUID 	`gorm:"type:uuid;not null"`
	IsActive 	bool 		`gorm:"type:boolean;not null"`
	CreatedBy 	string 		`gorm:"type:varchar(255);not null; default:'system'"`
	UpdatedBy 	string 		`gorm:"type:varchar(255);not null; default:'system'"`
	CreatedAt	time.Time 	`gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt 	time.Time 	`gorm:"autoUpdateTime;not null;default:now()"`

	// Start of References
	CartInformation []CartInformations `gorm:"foreignKey:CartID"`
	// End of References
}