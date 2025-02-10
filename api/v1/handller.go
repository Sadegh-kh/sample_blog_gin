package v1

import (
	"blog/api/v1/user"
)

type V1Handller struct {
	UserHandller user.UserHandller
}

func New(authS user.Service) V1Handller {
	return V1Handller{UserHandller: user.New(authS)}
}
