package container

import (
	//"learning-backend/config"
	"learning-backend/helper"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)
// A central container of dependencies 
// Instead of passing things again and again, bundle everything in one place


type HttpHandler struct {
	App *fiber.App
	DB *gorm.DB
	Auth helper.AuthHelper

	// Later
	//Config config.AppConfig
	//Pc payment.PaymentClient
}