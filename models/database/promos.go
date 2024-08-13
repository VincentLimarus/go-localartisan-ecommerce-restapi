package database

import (
	"time"

	"github.com/google/uuid"
)

type Promos struct {
	ID            	uuid.UUID `gorm:"type:uuid;primary_key;not null;default:uuid_generate_v4()"`
	ProductID    	uuid.UUID `gorm:"type:uuid;not null"`
	Name 			string    `gorm:"type:varchar(255);not null"`
	Description  	string    `gorm:"type:text;not null"`
	PromoDiscount 	float64   `gorm:"type:decimal(10,2);not null"`
	IsActive      	bool      `gorm:"type:boolean;not null;default:true"`
	CreatedBy     	string    `gorm:"type:varchar(255);not null; default:'system'"`
	UpdatedBy     	string    `gorm:"type:varchar(255);not null; default:'system'"`
	CreatedAt     	time.Time `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt     	time.Time `gorm:"type:timestamp;not null;default:now()"`
}
