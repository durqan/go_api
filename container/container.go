package container

import (
	"test/handlers"
	"test/middleware"
	"test/repository"
	"test/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Container struct {
	UserHandler *handlers.UserHandler
	Middleware  *MiddlewareContainer
}

type MiddlewareContainer struct {
	JWTAuth gin.HandlerFunc
}

func NewContainer(db *gorm.DB) *Container {
	userRepo := repository.NewUserRepository(db)
	passportRepo := repository.NewPassportRepository(db)
	addressRepo := repository.NewAddressRepository(db)
	jwtService := service.NewJWTService("123")

	return &Container{
		UserHandler: handlers.NewUserHandler(userRepo, passportRepo, addressRepo, jwtService),
		Middleware: &MiddlewareContainer{
			JWTAuth: middleware.JWTAuth(jwtService, userRepo),
		},
	}
}
