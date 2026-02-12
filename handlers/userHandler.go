package handlers

import (
	"learning-backend/database"
	dto_ "learning-backend/dto"
	//"learning-backend/dto"
	"learning-backend/models"

	"github.com/gofiber/fiber/v2"
)


func CreateUser(c *fiber.Ctx) error {
	var input dto_.CreateUserDTO  //dto.CreateUserDTO  

	if err := c.BodyParser(&input); err != nil {

		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})	
	}
  //Create model from DTO input
		user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password, // (later we’ll hash it)
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
		Name:  user.Name,
		Email: user.Email,
	}
	return c.Status(201).JSON(response)
}