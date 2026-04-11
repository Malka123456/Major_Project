package helper

import (
	"errors"
	"fmt"
	"learning-backend/config"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthHelper struct {

	Secret string
}


func (h AuthHelper) GenerateHashedPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", errors.New("Could not hash password")
	}
	return string(hashedPassword), nil
}

func (h AuthHelper) VerifyPassword(pP string, hP string) error {

	if len(pP) < 6 {
		return errors.New("password length should be at least 6 characters long")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hP), []byte(pP))

	if err != nil {
		return errors.New("password does not match")
	}

	return nil
}



func (h AuthHelper) GenerateToken(userID uint, email string, role string) (string, error) {

		if userID == 0 || email == "" || role == "" {
			return "", errors.New("invalid user data for token generation") 
		}	

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": userID,
			"email":   email,
			"role":    role,
			"exp":     jwt.TimeFunc().Add(24 * time.Hour * 7).Unix(), // Token expires in a week
		})

		tokenString, err := token.SignedString([]byte(h.Secret))

		if err != nil {
			return "", fmt.Errorf("unable to sign token: %w", err)
		}

		return tokenString, nil
}

func NewAuthHelper(cfg config.AppConfig) AuthHelper {
	return AuthHelper{
		Secret: cfg.JWTSecret,
		
	}
}