package models

import "time"

type Address struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint   `json:"user_id" gorm:"not null;index"`
	Full   string `json:"full" gorm:"type:varchar(255);not null"`
	Region string `json:"region" gorm:"type:varchar(50);not null"`
	City   string `json:"city" gorm:"type:varchar(50);not null"`
	Street string `json:"street" gorm:"type:varchar(50);not null"`
	House  string `json:"house" gorm:"type:varchar(4)"`
	Room   string `json:"room" gorm:"type:varchar(4)"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
