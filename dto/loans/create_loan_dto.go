package loans

import (
	"test/models"
)

type CreateLoanRequest struct {
	Sum  float64 `json:"sum" binding:"required,min=3000,max=30000"`
	Term int     `json:"term" binding:"required,min=1,max=28"`
}

func (r *CreateLoanRequest) ToLoanModel(userID uint) (models.Loans, error) {
	return models.Loans{
		UserID: userID,
		Sum:    r.Sum,
		Term:   r.Term,
	}, nil
}
