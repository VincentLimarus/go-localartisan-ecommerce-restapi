package database

import (
	"time"

	"github.com/google/uuid"
)

type CartInformations struct {
	CartID 			uuid.UUID 	`gorm:"type:uuid;not null;primaryKey"`
	ProductID 		uuid.UUID 	`gorm:"type:uuid;not null;primaryKey"`
	Quantity 		int 		`gorm:"type:int;not null"`
	PriceAtOrder 	float64 	`gorm:"type:float;not null"`
	IsActive 		bool 		`gorm:"type:boolean;not null"`
	CreatedBy 		string 		`gorm:"type:varchar(50);not null; default:'system'"`
	UpdatedBy 		string 		`gorm:"type:varchar(50);not null; default:'system'"`
	CreatedAt 		time.Time 	`gorm:"type:timestamp;not null; default:now()"`
	UpdatedAt 		time.Time 	`gorm:"type:timestamp;not null; default:now()"`
}