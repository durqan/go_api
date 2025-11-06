package handlers

import (
	"net/http"
	"test/dto/passport"
	"test/dto/user"
	"test/models"
	"test/repository"
	"test/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo     *repository.UserRepository
	passportRepo *repository.PassportRepository
	jwtService   *service.JWTService
}

func NewUserHandler(
	userRepo *repository.UserRepository,
	passportRepo *repository.PassportRepository,
	jwtService *service.JWTService,
) *UserHandler {
	return &UserHandler{
		userRepo:     userRepo,
		passportRepo: passportRepo,
		jwtService:   jwtService,
	}
}

func (h *UserHandler) AddUserWithContacts(c *gin.Context) {
	var request user.CreateUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверные данные: " + err.Error(),
		})
		return
	}

	userModel, err := request.ToUserModel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверный формат данных: " + err.Error(),
		})
		return
	}

	if err := h.userRepo.Create(&userModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при создании контакта: " + err.Error(),
		})
		return
	}

	token, err := h.jwtService.GenerateToken(userModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при генерации токена: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь успешно создан и авторизован",
		"token":   token,
		"user":    userModel,
	})
}

func (h *UserHandler) AddPassport(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	var request passport.CreatePassportRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	passportModel, err := request.ToPassportModel(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.passportRepo.Create(&passportModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании паспорта"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Паспортные данные успешно добавлены",
		"passport": passportModel,
	})
}

func (h *UserHandler) AddAddresses(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Contacts added successfully",
		"data":    "Users list",
	})
}
