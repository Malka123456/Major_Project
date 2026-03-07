package middleware

import (
	//"go/token"
	"learning-backend/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	//"github.com/golang-jwt/jwt/v4"
)


func AuthMiddleware(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Missing token",
		})	
	}

	tokenString := strings.Split(authHeader, " ")[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error)  {
		return config.PublicKey, nil
		})

		if err != nil || !token.Valid { 
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		return c.Next()

}



