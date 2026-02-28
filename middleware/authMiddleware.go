package middleware

import (
	//"go/token"
	"errors"
	"learning-backend/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
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

func GenerateHashedPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", errors.New("Could not hash password")
	}
	return string(hashedPassword), nil
}

func  VerifyPassword(pP string, hP string) error {

	if len(pP) < 6 {
		return errors.New("password length should be at least 6 characters long")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hP), []byte(pP))

	if err != nil {
		return errors.New("password does not match")
	}

	return nil
}
