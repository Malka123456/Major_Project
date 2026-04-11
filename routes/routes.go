package routes

import (
	"learning-backend/handlers"
	//"learning-backend/middleware"

	"github.com/gofiber/fiber/v2"
)
type RouteHandler struct {
	handlers handlers.UserHandler
}



func (r *RouteHandler) SetupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Backend is running")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Server is healthy")
	})
	
	pubRoutes := app.Group("/api")
	
	pubRoutes.Post("/signup", r.handlers.SignUp) 
	pubRoutes.Post("/signin", r.handlers.SignIn)

	// priRoutes := app.Group("/user", middleware.AuthMiddleware) // Apply auth middleware to all routes in this group

	// priRoutes.Get("/profile", handlers.GetProfile)	

	// app.Get("/profile", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 				"message": "Welcome",
	// 			})
	// 		},)
		
	


}

// Constructor (BEST PRACTICE)
func NewRouteHandler(h handlers.UserHandler) RouteHandler {
    return RouteHandler{
        handlers: h,
    }
}
