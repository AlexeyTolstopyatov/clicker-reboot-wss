package userHost

import (
	"clicker/internal/app/dbhost"
	"clicker/internal/app/txhost"
)

type UserHost struct {
	userApi  dbhost.UserQueryable
	teamApi  dbhost.TeamQueryable
	thostApi txhost.Transactable
}

func Instance(userApi dbhost.UserQueryable, teamApi dbhost.TeamQueryable, txh txhost.Transactable) *UserHost {
	return &UserHost{userApi, teamApi, txh}
}

// Other UserHost API implemented in UserImpl.go
