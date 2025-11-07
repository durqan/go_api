package models

import "time"

type Loans struct {
	ID     uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint    `json:"user_id" gorm:"not null;index"`
	User   User    `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Sum    float64 `json:"amount" gorm:"type:decimal(10,2);not null;index"`
	Term   int     `json:"term" gorm:"type:integer;not null;index"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
