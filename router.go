package main

import (
	"test/container"
	"test/database"

	"github.com/gin-gonic/gin"
)

func startRouter() {
	r := gin.Default()
	db := database.GetDB()

	container := container.NewContainer(db)

	api := r.Group("/api")
	{
		api.POST("/add_contacts", container.UserHandler.AddUserWithContacts)

		api.Use(container.Middleware.JWTAuth)
		{
			api.POST("/add_passport", container.UserHandler.AddPassport)
			api.POST("/add_addresses", container.UserHandler.AddAddresses)
		}
	}

	r.Run(":8082")
}
