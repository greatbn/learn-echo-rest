package models

type (
	User struct {
		Name     string `json:"name" validate:"nonzero"`
		Password string `json:"password" validate:"nonzero"`
		Email    string `json:"email" validate:"nonzero"`
	}
)
