package database

import (
	"time"

	"github.com/google/uuid"
)

type CartInformations struct {
	CartID 			uuid.UUID 	`gorm:"type:uuid;primaryKey;not null"`
	ProductID 		uuid.UUID 	`gorm:"type:uuid;primaryKey;not null"`
	Quantity 		int 		`gorm:"type:int;not null"`
	PriceAtOrder 	float64 	`gorm:"type:float;not null"`
	IsActive 		bool 		`gorm:"type:boolean;not null"`
	CreatedBy 		string 		`gorm:"type:varchar(255);not null"`
	UpdatedBy 		string 		`gorm:"type:varchar(255);not null"`
	CreatedAt 		time.Time 	`gorm:"type:timestamp;not null"`
	UpdatedAt 		time.Time 	`gorm:"type:timestamp;not null"`
}