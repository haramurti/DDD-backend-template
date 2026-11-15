package contract

import (
	"context"

	"hackathon-be/internal/app/user/dto"
)

type UserService interface {
	Register(ctx context.Context, req dto.RegisterRequest) (*dto.UserResponse, error)
	Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error)
	Me(ctx context.Context, userID uint) (*dto.UserResponse, error)
}
