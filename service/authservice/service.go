package authservice

import (
	"blog/api/form"
)

type storage interface {
	Register(user form.UserRegister) error
	CheckUsername(userName string) error
	CheckEmail(email string) error
}

type Service struct {
	Storage storage
}
