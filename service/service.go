package service

import "blog/service/authservice"

type Service struct {
	AuthService authservice.Service
}

func New(s authservice.Storage) Service {
	service := Service{AuthService: authservice.New(s)}

	return service
}
