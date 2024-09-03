package database

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatusType string

const (
	WaitingForPayment OrderStatusType = "Waiting for Payment"
	OrderFinished     OrderStatusType = "Order Finished"
	OrderCanceled     OrderStatusType = "Order Canceled"
	OrderPaid 	   	  OrderStatusType = "Order Paid"
)

type Orders struct {
	ID               uuid.UUID        `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID           uuid.UUID        `gorm:"type:uuid;not null"`
	Status           OrderStatusType  `gorm:"type:enum('Waiting for Payment', 'Order Finished', 'Order Canceled', 'Order Paid');not null;default:'Waiting for Payment'"`
	TotalPrice       float64          `gorm:"type:float;not null"`
	ShippingAddress  string           `gorm:"type:text;not null"`
	PaymentMethod    string           `gorm:"type:varchar(255);not null"`
	IsActive         bool             `gorm:"type:boolean;not null"`
	CreatedBy        string           `gorm:"type:varchar(255);not null;default:'system'"`
	UpdatedBy        string           `gorm:"type:varchar(255);not null;default:'system'"`
	CreatedAt        time.Time        `gorm:"type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt        time.Time        `gorm:"type:timestamp with time zone;not null;default:CURRENT_TIMESTAMP"`

	OrderItems       []OrderItems     `gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
