package service

import (
	"context"

	"hackathon-be/internal/app/user/contract"
	"hackathon-be/internal/app/user/dto"
	"hackathon-be/pkg/response"
)

type service struct {
	repo      contract.UserRepository
	jwtSecret string
}

func New(repo contract.UserRepository, jwtSecret string) contract.UserService {
	return &service{repo: repo, jwtSecret: jwtSecret}
}

func (s *service) Register(ctx context.Context, req dto.RegisterRequest) (*dto.UserResponse, error) {
	// TODO: lowercase email -> utils.HashPassword -> repo.Create (role default member)
	// duplikat email -> response.ErrConflict
	return nil, response.ErrNotImplemented
}

func (s *service) Login(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	// TODO: repo.FindByEmail -> utils.CheckPassword -> jwt.Sign(s.jwtSecret, u.ID, u.Role, 24*time.Hour)
	// email/password salah -> response.ErrUnauthorized (pesan sama, jangan bocorin email terdaftar)
	return nil, response.ErrNotImplemented
}

func (s *service) Me(ctx context.Context, userID uint) (*dto.UserResponse, error) {
	// TODO
	return nil, response.ErrNotImplemented
}
