package main

import (
	"learning-backend/database"
	//"learning-backend/handlers"
	//"learning-backend/models"

	"fmt"

	"github.com/gofiber/fiber/v2"
)


func main() {

	database.ConnectDB()

	app := fiber.New()

	
	fmt.Println("server is running")

	app.Listen(":3000")

	
}