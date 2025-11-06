package repository

import (
	"test/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User

	result := r.db.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
