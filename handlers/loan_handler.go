package handlers

import (
	"net/http"
	"test/dto/loans"
	"test/models"
	"test/repository"
	"test/service"

	"github.com/gin-gonic/gin"
)

type LoanHandler struct {
	userRepo   *repository.UserRepository
	loanRepo   *repository.LoanRepository
	jwtService *service.JWTService
}

func NewLoanHandler(
	userRepo *repository.UserRepository,
	loanRepo *repository.LoanRepository,
	jwtService *service.JWTService,
) *LoanHandler {
	return &LoanHandler{
		userRepo:   userRepo,
		loanRepo:   loanRepo,
		jwtService: jwtService,
	}
}

// CreateLoan godoc
// @Summary Add loan
// @Description Create a new loan
// @Tags loans
// @Accept json
// @Produce json
// @Param request body loans.CreateLoanRequest true "Loan with data"
// @Success 201 {object} object "Loan created successfully"
// @Failure 400 {object} object "Invalid request data"
// @Failure 500 {object} object "Internal server error"
// @Router /add_loan [post]
func (h *LoanHandler) CreateLoan(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	var request loans.CreateLoanRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	loanModel, err := request.ToLoanModel(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.loanRepo.Create(&loanModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании паспорта"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Паспортные данные успешно добавлены",
		"passport": loanModel,
	})
}
