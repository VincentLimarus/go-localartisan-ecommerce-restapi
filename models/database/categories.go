package database

import (
	"time"

	"github.com/google/uuid"
)

type Categories struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4(); not null"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Image     string    `gorm:"type:varchar(255);"`
	IsActive  bool      `gorm:"type:boolean;not null;default:true"`
	CreatedBy string    `gorm:"type:varchar(50);not null; default:'system'"`
	UpdatedBy string    `gorm:"type:varchar(50);not null; default:'system'"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null;default:now()"`

	// Start of References
	Product []Product `gorm:"foreignKey:CategoryID"`
	// End of References
}