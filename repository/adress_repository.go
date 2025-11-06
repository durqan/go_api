package repository

import (
	"test/models"

	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{db: db}
}

func (r *AddressRepository) Create(address *models.Address) error {
	return r.db.Create(address).Error
}
