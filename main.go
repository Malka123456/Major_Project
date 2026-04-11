package main

import (
	"learning-backend/config"
	"learning-backend/container"
	"learning-backend/routes"

	//"learning-backend/handlers"
	//"learning-backend/models"

	"fmt"

	"github.com/gofiber/fiber/v2"
)




func main() {


	app := fiber.New()

	container := container.BuildContainer()

	routes.SetupRoutes(app, container)

	config.LoadKeys()


	fmt.Println("server is running")

	app.Listen(":3000")
}