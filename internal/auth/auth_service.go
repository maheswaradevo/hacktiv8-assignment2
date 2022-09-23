package auth

import (
	"context"
	"database/sql"

	"github.com/maheswaradevo/hacktiv8-assignment2/internal/auth/impl"
	"github.com/maheswaradevo/hacktiv8-assignment2/internal/dto"
)

type AuthService interface {
	NewUser(ctx context.Context, data *dto.UserRegistrationRequest) (userID int64, err error)
	LoginUser(ctx context.Context, data *dto.UserLoginRequest) (res *dto.UserLoginResponse, err error)
}

func ProvideAuthService(DB *sql.DB) AuthService {
	repo := impl.ProvideAuthRepository(DB)
	return impl.ProvideAuthService(repo)
}
