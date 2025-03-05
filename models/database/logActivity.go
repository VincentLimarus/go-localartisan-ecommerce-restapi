package database

import "time"

type LogActivity struct {
	ID        		uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Action    		string    `json:"action"`
	UpdatedBy    	string    `json:"updated_by"`
	UpdatedAt 		time.Time `gorm:"autoCreateTime" json:"updated_at"`
}
