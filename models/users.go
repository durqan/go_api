package models

import "time"

type User struct {
	ID         uint          `json:"id" gorm:"primaryKey;autoIncrement"`
	Lastname   string        `json:"lastname" gorm:"type:varchar(100);not null;index"`
	Firstname  string        `json:"firstname" gorm:"type:varchar(100);not null;index"`
	Patronymic string        `json:"patronymic,omitempty" gorm:"type:varchar(100);default:null;index"`
	BirthDate  time.Time     `json:"birth_date" gorm:"type:date;index"`
	Sex        string        `json:"sex" gorm:"type:varchar(10);check:sex IN ('male', 'female');index"`
	Email      string        `json:"email" gorm:"type:varchar(100);uniqueIndex"`
	Phone      string        `json:"phone" gorm:"type:varchar(100);uniqueIndex"`
	Passport   Passport      `json:"passport,omitempty" gorm:"foreignKey:UserID"`
	Address    []Address     `json:"reg_address,omitempty" gorm:"foreignKey:UserID"`
	Contacts   []UserContact `json:"contacts" gorm:"foreignKey:UserID"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
