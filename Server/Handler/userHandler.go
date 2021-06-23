package handler

import (
	"assessment-bootcamp/auth"
	"os/user"
)

type userhandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service)
