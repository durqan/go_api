package service

import (
	"errors"
	"test/dto/contact"
	"test/models"

	"gorm.io/gorm"
)

type ContactService struct {
	db *gorm.DB
}

func (s *ContactService) CreateContact(userID uint, req contact.CreateContactRequest) (*models.UserContact, error) {
	if exists, _ := s.ContactExists(userID, req.Type, req.Value); exists {
		return nil, errors.New("контакт уже существует")
	}

	contact := &models.UserContact{
		UserID: userID,
		Type:   req.Type,
		Value:  req.Value,
	}

	if err := s.db.Create(contact).Error; err != nil {
		return nil, err
	}

	return contact, nil
}

func (s *ContactService) ContactExists(userID uint, contactType, value string) (bool, error) {
	var count int64
	err := s.db.Model(&models.UserContact{}).
		Where("user_id = ? AND type = ? AND value = ?", userID, contactType, value).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
