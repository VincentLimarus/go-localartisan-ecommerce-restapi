package database

import (
	"time"

	"github.com/google/uuid"
)

type Carts struct {
	ID 			uuid.UUID 	`gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID 		uuid.UUID 	`gorm:"type:uuid;not null"`
	IsActive 	bool 		`gorm:"type:boolean;not null"`
	CreatedBy 	string 		`gorm:"type:varchar(255);not null"`
	UpdatedBy 	string 		`gorm:"type:varchar(255);not null"`
	CreatedAt	time.Time 	`gorm:"type:timestamp;not null"`
	UpdatedAt 	time.Time 	`gorm:"type:timestamp;not null"`
}