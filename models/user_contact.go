package models

import "time"

type UserContact struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID uint   `json:"user_id" gorm:"not null;index"`
	Type   string `json:"type" gorm:"type:varchar(50);not null;index"`
	Value  string `json:"value" gorm:"type:varchar(255);not null"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

const (
	ContactTypeEmail    = "email"
	ContactTypePhone    = "phone"
	ContactTypeWork     = "work"
	ContactTypeTelegram = "telegram"
	ContactTypeWhatsApp = "whatsapp"
)

func (uc *UserContact) IsValidType() bool {
	validTypes := map[string]bool{
		ContactTypeEmail:    true,
		ContactTypePhone:    true,
		ContactTypeWork:     true,
		ContactTypeTelegram: true,
		ContactTypeWhatsApp: true,
	}
	return validTypes[uc.Type]
}
