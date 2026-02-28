package handlers

import (
	//"learning-backend/config"
	//"hash"
	"learning-backend/database"
	dto_ "learning-backend/dto"
	"learning-backend/middleware"

	//"learning-backend/dto"
	"learning-backend/models"

	"github.com/gofiber/fiber/v2"
	//"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	var input dto_.CreateUserDTO //dto.CreateUserDTO

	if err := c.BodyParser(&input); err != nil {

		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}
	//code to prevent duplicate emails record
	var existingUser models.User
	database.DB.Where("email=?", input.Email).First(&existingUser)
	if existingUser.ID != 0 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Email already exists",
		})

	}

	hashedPassword, err := middleware.GenerateHashedPassword(input.Password)

	if err != nil {
		return err
	}

	//Create model from DTO input
	user := models.User{
		FirstName:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword), // (later we’ll hash it)
	}

	//save to database
	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": result.Error.Error(),
		})
	}

	// Send response DTO (hide password)
	response := dto_.UserResponseDTO{
		ID:    user.ID,
		Name:  user.FirstName,
		Email: user.Email,
	}
	return c.Status(201).JSON(response)
}

func Login(c *fiber.Ctx) error {
	var input dto_.LoginUserDTO

	if err := c.BodyParser(&input); err != nil {

		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	password, err := middleware.GenerateHashedPassword(input.Password)
  if err != nil {
		return err
	}

	err = middleware.VerifyPassword(input.Password, password)
	if err != nil {
		return err
	}

	return nil

}
