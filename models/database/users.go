package database

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID 			uuid.UUID 	`gorm:"primaryKey;type:uuid;default:uuid_generate_v4(); not null"`
	Name 		string 		`gorm:"type:varchar(255);not null"`
	Email 		string 		`gorm:"type:varchar(255);not null;unique"`
	Password 	string 		`gorm:"type:varchar(255);not null"`
	PhoneNumber string 		`gorm:"type:varchar(255);"`
	IsActive 	bool 		`gorm:"type:boolean;default:true"`
	CreatedBy 	string 		`gorm:"type:varchar(50);not null" default:"system"`
	UpdatedBy 	string 		`gorm:"type:varchar(50);not null" default:"system"`
	CreatedAt 	time.Time 	`gorm:"type:autoCreateTime;not null" default:now()"`
	UpdatedAt 	time.Time 	`gorm:"type:autoUpdateTime;not null" default:now()"`
}
