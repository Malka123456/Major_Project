package container

import (
	"learning-backend/config"
	"learning-backend/database"
	"learning-backend/handlers"
	"learning-backend/helper"
	"learning-backend/repository"
	"learning-backend/service"

	"gorm.io/gorm"
)

type Container struct {
	// 🔹 Core dependencies
	DB     *gorm.DB
	Config config.AppConfig
	Auth   helper.AuthHelper

	// 🔹 Services
	UserService *service.UserService

	// 🔹 Handlers
	UserHandler *handlers.UserHandler
}

func BuildContainer() *Container {

	// 🔹 base
	config := config.LoadConfig()

	db := database.InitDB(config.DBUrl)
	auth := helper.NewAuthHelper(config)

	// 🔹 repositories
	userRepo := repository.NewUserRepository(db)

	// 🔹 services
	userService := service.NewUserService(userRepo, auth)

	// 🔹 handlers
	userHandler := handlers.NewUserHandler(userService)

	return &Container{
		DB: db,
		Config: config,
		Auth: auth,

		UserService: userService,
		UserHandler: userHandler,
	}
}