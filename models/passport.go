package models

import "time"

type Passport struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID          uint      `json:"user_id" gorm:"not null;uniqueIndex"`
	Series          string    `json:"series" gorm:"type:varchar(10);not null"`
	Number          string    `json:"number" gorm:"type:varchar(20);not null"`
	IssueDate       time.Time `json:"issue_date" gorm:"type:date;not null"`
	IssueDepartment string    `json:"issue_department" gorm:"type:varchar(500);not null"`
	BirthPlace      string    `json:"birth_place" gorm:"type:varchar(200)"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
