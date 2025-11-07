package handlers

import (
	"net/http"
	"test/dto/address"
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
	addressRepo  *repository.AddressRepository
	jwtService   *service.JWTService
}

func NewUserHandler(
	userRepo *repository.UserRepository,
	passportRepo *repository.PassportRepository,
	addressRepo *repository.AddressRepository,
	jwtService *service.JWTService,
) *UserHandler {
	return &UserHandler{
		userRepo:     userRepo,
		passportRepo: passportRepo,
		addressRepo:  addressRepo,
		jwtService:   jwtService,
	}
}

// AddUserWithContacts godoc
// @Summary Add user with contacts
// @Description Create a new user with contact information
// @Tags users
// @Accept json
// @Produce json
// @Param request body user.CreateUserRequest true "User with contacts data"
// @Success 201 {object} object "User created successfully"
// @Failure 400 {object} object "Invalid request data"
// @Failure 500 {object} object "Internal server error"
// @Router /add_contacts [post]
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

// AddPassport godoc
// @Summary Add passport information
// @Description Add passport details for authenticated user
// @Tags passport
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body passport.CreatePassportRequest true "Passport data"
// @Success 201 {object} object "Passport added successfully"
// @Failure 400 {object} object "Invalid request data"
// @Failure 401 {object} object "Unauthorized"
// @Failure 500 {object} object "Internal server error"
// @Router /add_passport [post]
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

// AddAddresses godoc
// @Summary Add addresses
// @Description Add address information for authenticated user
// @Tags addresses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body address.CreateAddressRequest true "Addresses data"
// @Success 201 {object} object "Address added successfully"
// @Failure 400 {object} object "Invalid request data"
// @Failure 401 {object} object "Unauthorized"
// @Failure 500 {object} object "Internal server error"
// @Router /add_addresses [post]
func (h *UserHandler) AddAddresses(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	var request address.CreateAddressRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные: " + err.Error()})
		return
	}

	addressModel := request.ToAddressModel(user.ID)

	if err := h.addressRepo.Create(&addressModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании адреса"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Адрес успешно добавлен",
		"address": addressModel,
	})
}
