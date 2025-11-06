package repository

import (
	"test/models"

	"gorm.io/gorm"
)

type PassportRepository struct {
	db *gorm.DB
}

func NewPassportRepository(db *gorm.DB) *PassportRepository {
	return &PassportRepository{db: db}
}

func (r *PassportRepository) Create(passport *models.Passport) error {
	return r.db.Create(passport).Error
}

func (r *PassportRepository) FindByUserID(userID uint) (*models.Passport, error) {
	var passport models.Passport
	err := r.db.Where("user_id = ?", userID).First(&passport).Error
	if err != nil {
		return nil, err
	}
	return &passport, nil
}
