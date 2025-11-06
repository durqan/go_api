package passport

import (
	"test/models"
	"time"
)

type CreatePassportRequest struct {
	Series          string `json:"Series" binding:"required,len=4"`
	Number          string `json:"Number" binding:"required,len=6"`
	IssueDate       string `json:"IssueDate" binding:"required"`
	IssueDepartment string `json:"IssueDepartment" binding:"required,min=5,max=500"`
	BirthPlace      string `json:"BirthPlace" binding:"max=200"`
}

func (r *CreatePassportRequest) ToPassportModel(userID uint) (models.Passport, error) {
	issueDate, err := time.Parse("2006-01-02", r.IssueDate)
	if err != nil {
		return models.Passport{}, err
	}

	return models.Passport{
		UserID:          userID,
		Series:          r.Series,
		Number:          r.Number,
		IssueDate:       issueDate,
		IssueDepartment: r.IssueDepartment,
		BirthPlace:      r.BirthPlace,
	}, nil
}
