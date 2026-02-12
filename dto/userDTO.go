package dto_

type CreateUserDTO struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserResponseDTO struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}