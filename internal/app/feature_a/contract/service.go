package contract

import (
	"context"

	"hackathon-be/internal/app/feature_a/dto"
)

type FeatureAService interface {
	Create(ctx context.Context, actorID uint, req dto.CreateFeatureARequest) (*dto.FeatureAResponse, error)
	GetByID(ctx context.Context, actorID, id uint) (*dto.FeatureAResponse, error)
	List(ctx context.Context, actorID uint, q dto.ListFeatureAQuery) ([]dto.FeatureAResponse, int64, error)
	Update(ctx context.Context, actorID, id uint, req dto.UpdateFeatureARequest) (*dto.FeatureAResponse, error)
	Delete(ctx context.Context, actorID, id uint) error
}
