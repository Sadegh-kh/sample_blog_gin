package user

import "blog/api/form"

type Service interface {
	Register(user form.UserRegister) error
	ValidationRegister(user form.UserRegister) error
}

type UserHandller struct {
	AuthService Service
}

func New(AuthS Service) UserHandller {
	return UserHandller{AuthService: AuthS}
}
