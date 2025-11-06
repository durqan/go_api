package address

import (
	"test/models"
)

type CreateAddressRequest struct {
	Full   string `json:"Full" binding:"required"`
	Region string `json:"Region" binding:"required"`
	City   string `json:"City" binding:"required"`
	Street string `json:"Street" binding:"required"`
	House  string `json:"House" binding:"required"`
	Room   string `json:"Room"`
}

func (r *CreateAddressRequest) ToAddressModel(userID uint) models.Address {
	return models.Address{
		UserID: userID,
		Full:   r.Full,
		Region: r.Region,
		City:   r.City,
		Street: r.Street,
		House:  r.House,
		Room:   r.Room,
	}
}
