package dto_

type SignUp struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Phone string `json:"phone" validate:"required"`
}

type SignIn struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`

}

