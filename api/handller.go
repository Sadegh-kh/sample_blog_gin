package api

import (
	v1 "blog/api/v1"
	"blog/api/v1/user"
)

type ApiHandller struct {
	V1Handller v1.V1Handller
}

func New(authS user.Service) ApiHandller {
	return ApiHandller{V1Handller: v1.New(authS)}
}
