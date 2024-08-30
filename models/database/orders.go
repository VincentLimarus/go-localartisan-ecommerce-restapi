package database

import (
	"time"

	"github.com/google/uuid"
)

type Orders struct {
	ID 				uuid.UUID 		`gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID 			uuid.UUID 		`gorm:"type:uuid;not null"`
	Status 			string 			`gorm:"type:varchar(255);not null"`
	TotalPrice 		float64 		`gorm:"type:float;not null"`
	ShippingAddress string 			`gorm:"type:text;not null"`
	PaymentMethod 	string 			`gorm:"type:varchar(255);not null"`
	IsActive 		bool 			`gorm:"type:boolean;not null"`
	CreatedBy  	    string     `gorm:"type:varchar(255);not null; default:'system'"`
	UpdatedBy   	string     `gorm:"type:varchar(255);not null; default:'system'"`
	CreatedAt  	    time.Time  `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt   	time.Time  `gorm:"type:timestamp;not null;default:now()"`

	OrderItems 		[]OrderItems 	`gorm:"foreignKey:OrderID;references:ID"`
}