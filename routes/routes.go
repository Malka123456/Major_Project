package routes

import (
	"learning-backend/container"
	//"learning-backend/middleware"

	"github.com/gofiber/fiber/v2"
)




func  SetupRoutes(app *fiber.App, c *container.Container) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Backend is running")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Server is healthy")
	})
	
	pubRoutes := app.Group("/api")
	
	pubRoutes.Post("/signup", c.UserHandler.SignUp) 
	pubRoutes.Post("/signin", c.UserHandler.SignIn)

	// priRoutes := app.Group("/user", middleware.AuthMiddleware) // Apply auth middleware to all routes in this group

	// priRoutes.Get("/profile", handlers.GetProfile)	

	// app.Get("/profile", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 				"message": "Welcome",
	// 			})
	// 		},)
		
	


}

