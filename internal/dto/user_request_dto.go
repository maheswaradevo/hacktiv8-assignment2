package dto

import "github.com/maheswaradevo/hacktiv8-assignment2/internal/entity"

type UserRegistrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ur *UserRegistrationRequest) ToEntity() (u *entity.User) {
	u = &entity.User{
		Username: ur.Username,
		Email:    ur.Email,
		Password: ur.Password,
	}
	return
}

func (ul *UserLoginRequest) ToEntity() (us *entity.User) {
	us = &entity.User{
		Username: ul.Username,
		Password: ul.Password,
	}
	return
}
