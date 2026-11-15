package dto

import (
	"time"

	"hackathon-be/internal/app/feature_b/entity"
)

type FeatureBResponse struct {
	ID        uint      `json:"id"`
	OwnerID   uint      `json:"owner_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// TODO: field lain yang boleh dilihat client (jangan bocorin data privat)
}

func FromFeatureB(e entity.FeatureB) FeatureBResponse {
	return FeatureBResponse{
		ID:        e.ID,
		OwnerID:   e.OwnerID,
		Status:    e.Status,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
