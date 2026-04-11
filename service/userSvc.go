package service

import (
	"errors"
	dto_ "learning-backend/dto"
	"learning-backend/container"
	"learning-backend/helper"
	"learning-backend/models"
	"learning-backend/repository"
)



type UserService struct {
	Auth helper.AuthHelper
	Repo repository.UserRepository
}	


func (s *UserService) SignUp(input dto_.SignUp) (string, error) {

	hashedPassword, err := s.Auth.GenerateHashedPassword(input.Password) //dependency injected into service

	if err != nil {
		return "", err
	}

	//Create model from DTO input
	user, err := s.Repo.CreateUser(models.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone: input.Phone, 
	})	

	if err != nil {
		return "", err
	}

	
	return s.Auth.GenerateToken(user.ID, user.Email, string(user.UserType)) //generate token after creating user
}

func (s *UserService) findUserByEmail(email string) (*models.User, error) {
	
	user, err := s.Repo.FindUser(email)

	return &user, err
}

func (s *UserService) SignIn(email string, password string) (string, error) {

	user, err := s.findUserByEmail(email)
	if err != nil {
		return "", errors.New("user does not exist with the provided email id")
	}

	err = s.Auth.VerifyPassword(password, user.Password)

	if err != nil {
		return "", err
	}

	// generate token
	return s.Auth.GenerateToken(user.ID, user.Email, string(user.UserType))
}

func NewUserService(h container.HttpHandler) UserService {
	return UserService{
		Auth: h.Auth,
		Repo: repository.NewUserRepository(h.DB),
	}
}