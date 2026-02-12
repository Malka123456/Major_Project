package main

import (
	"learning-backend/database"
	"learning-backend/handlers"
	//"learning-backend/models"

	"fmt"

	"github.com/gofiber/fiber/v2"
)


func main() {

	database.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Backend is running")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Server is healthy")
	})
	
  app.Post("/register", handlers.CreateUser)


	fmt.Println("server is running")

	app.Listen(":3000")

	
}