package user

import (
	"test/models"
	"time"
)

type CreateUserRequest struct {
	Lastname   string `json:"Lastname" binding:"required,min=1,max=100"`
	Firstname  string `json:"Firstname" binding:"required,min=1,max=100"`
	Patronymic string `json:"Patronymic,omitempty" binding:"max=100"`
	BirthDate  string `json:"BirthDate" binding:"required"`
	Sex        string `json:"Sex" binding:"required,oneof=male female"`
	Email      string `json:"Email" binding:"required,email"`
	Phone      string `json:"Phone" binding:"required"`
}

func (r *CreateUserRequest) ToUserModel() (models.User, error) {

	birthDate, err := time.Parse("2006-01-02", r.BirthDate)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		Lastname:   r.Lastname,
		Firstname:  r.Firstname,
		Patronymic: r.Patronymic,
		BirthDate:  birthDate,
		Sex:        r.Sex,
		Email:      r.Email,
		Phone:      r.Phone,
	}, nil
}
