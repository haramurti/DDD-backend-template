package contract

import (
	"context"

	"hackathon-be/internal/app/feature_b/dto"
)

type FeatureBService interface {
	Create(ctx context.Context, actorID uint, req dto.CreateFeatureBRequest) (*dto.FeatureBResponse, error)
	GetByID(ctx context.Context, actorID, id uint) (*dto.FeatureBResponse, error)
	List(ctx context.Context, actorID uint, q dto.ListFeatureBQuery) ([]dto.FeatureBResponse, int64, error)
	Update(ctx context.Context, actorID, id uint, req dto.UpdateFeatureBRequest) (*dto.FeatureBResponse, error)
	Delete(ctx context.Context, actorID, id uint) error
}
