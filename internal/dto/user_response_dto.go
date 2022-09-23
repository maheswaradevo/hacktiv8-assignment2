package dto

import "github.com/maheswaradevo/hacktiv8-assignment2/internal/entity"

type UserLoginResponse struct {
	Username    string `json:"username"`
	AccessToken string `json:"access_token"`
}

func NewUserLoginResponse(user *entity.User, ac string) (res *UserLoginResponse, err error) {
	res = &UserLoginResponse{
		Username:    user.Username,
		AccessToken: ac,
	}
	return
}
