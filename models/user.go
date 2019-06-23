package models

import "github.com/go-playground/validator"

type (
	User struct {
		Name     string `json:"name" validate:"required"`
		Password string `json:"password" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
	}

	UserValidator struct {
		Validator *validator.Validate
	}
)

func (uv *UserValidator) Validate(i interface{}) error {
	return uv.Validator.Struct(i)
}
