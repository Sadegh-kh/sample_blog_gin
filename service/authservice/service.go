package authservice

import (
	"blog/api/form"
)

type Storage interface {
	Register(user form.UserRegister) error
	CheckUniqUsername(userName string) (bool, error)
	CheckUniqEmail(email string) (bool, error)
}

type Service struct {
	Storage Storage
}

func New(s Storage) Service {
	return Service{Storage: s}
}
