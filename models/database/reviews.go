package database

import (
	"time"

	"github.com/google/uuid"
)

type Reviews struct {
	ID 			uuid.UUID 	`gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProductID 	uuid.UUID 	`gorm:"type:uuid;not null"`
	UserID 		uuid.UUID 	`gorm:"type:uuid;not null"`
	Comment 	string 		`gorm:"type:text;not null"`
	Rating 		float64 	`gorm:"type:float;not null"`
	IsActive 	bool 		`gorm:"type:boolean;not null;default:true"`
	CreatedBy 	string	 	`gorm:"type:varchar(255);not null;default:system"`
	UpdatedBy 	string 		`gorm:"type:varchar(255);not null;default:system"`
	CreatedAt 	time.Time 	`gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt 	time.Time 	`gorm:"type:timestamp;not null;default:now()"`
}