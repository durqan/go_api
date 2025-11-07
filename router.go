// @title User API
// @version 1.0
// @description This is a sample user management API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8082
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT Authorization header using the Bearer scheme
package main

import (
	"test/container"
	"test/database"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func startRouter() {
	r := gin.Default()
	db := database.GetDB()

	container := container.NewContainer(db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		api.POST("/add_contacts", container.UserHandler.AddUserWithContacts)

		api.Use(container.Middleware.JWTAuth)
		{
			api.POST("/add_passport", container.UserHandler.AddPassport)
			api.POST("/add_addresses", container.UserHandler.AddAddresses)
			api.POST("/add_loan", container.LoanHandler.CreateLoan)
		}
	}

	r.Run(":8082")
}
