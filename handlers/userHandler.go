package handlers

import (
	//"learning-backend/config"
	//"hash"
	dto_ "learning-backend/dto"
	"learning-backend/service"
	"net/http"

	//"learning-backend/middleware"

	//"learning-backend/dto"

	"github.com/gofiber/fiber/v2"
	//"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	//svc *service.UserService
		Service *service.UserService
}



func (h *UserHandler) SignUp(c *fiber.Ctx) error {
	var input dto_.SignUp //dto.CreateUserDTO

	if err := c.BodyParser(&input); err != nil {

		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}
	

	
	token, err := h.Service.SignUp(input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error on signup",
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
		"token":   token,
	})

}

func (h *UserHandler) SignIn(c *fiber.Ctx) error {
	var input dto_.SignIn

	if err := c.BodyParser(&input); err != nil {

		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	token, err := h.Service.SignIn(input.Email, input.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "invalid credentials",
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login successful",
		"token":   token,
	})


}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}