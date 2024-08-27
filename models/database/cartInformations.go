package database

import (
	"time"

	"github.com/google/uuid"
)

type CartInformations struct {
	ID 				uuid.UUID 	`gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CartID 			uuid.UUID 	`gorm:"type:uuid;not null;"`
	ProductID 		uuid.UUID 	`gorm:"type:uuid;not null;"`
	Quantity 		int 		`gorm:"type:int;not null"`
	PriceAtOrder 	float64 	`gorm:"type:float;not null"`
	IsActive 		bool 		`gorm:"type:boolean;not null"`
	CreatedBy 		string 		`gorm:"type:varchar(50);not null; default:'system'"`
	UpdatedBy 		string 		`gorm:"type:varchar(50);not null; default:'system'"`
	CreatedAt 		time.Time 	`gorm:"type:timestamp;not null; default:now()"`
	UpdatedAt 		time.Time 	`gorm:"type:timestamp;not null; default:now()"`
}