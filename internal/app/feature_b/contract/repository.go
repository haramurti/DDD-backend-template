package contract

import (
	"context"

	"hackathon-be/internal/app/feature_b/dto"
	"hackathon-be/internal/app/feature_b/entity"

	"gorm.io/gorm"
)

type FeatureBRepository interface {
	WithTx(tx *gorm.DB) FeatureBRepository
	Create(ctx context.Context, e *entity.FeatureB) error
	FindByID(ctx context.Context, id uint) (*entity.FeatureB, error)
	List(ctx context.Context, q dto.ListFeatureBQuery) ([]entity.FeatureB, int64, error)
	Update(ctx context.Context, e *entity.FeatureB) error
	Delete(ctx context.Context, id uint) error
	// TODO: method khusus, contoh conditional update:
	// UpdateIfStatus(ctx, id, ownerID uint, fromStatus string, fields map[string]any) (rows int64, err error)
}
