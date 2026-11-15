package dto

import "hackathon-be/internal/app/user/entity"

type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
	// email sengaja gak ditampilkan; tambahin kalau PRD minta
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

func FromUser(u entity.User) UserResponse {
	return UserResponse{ID: u.ID, Name: u.Name, Role: u.Role}
}
