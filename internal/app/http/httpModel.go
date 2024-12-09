package http

import (
	"clicker/internal/app/userHost"
)

type Handler struct {
	UserApi userHost.UserHost
}

func Instance(userApi userHost.UserHost) *Handler {
	return &Handler{
		UserApi: userApi,
	}
}
