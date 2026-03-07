package main

import (
	"learning-backend/config"
	"learning-backend/database"
	"learning-backend/handlers"
	"learning-backend/routes"

	//"learning-backend/handlers"
	//"learning-backend/models"

	"fmt"

	"github.com/gofiber/fiber/v2"
)


func main() {

	database.ConnectDB()

	app := fiber.New()

	
	handlers.SetupRoutes(app)
	
	config.LoadKeys()

	
	fmt.Println("server is running")

	app.Listen(":3000")

	
}