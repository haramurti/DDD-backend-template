package contract

import (
	"context"

	"hackathon-be/internal/app/feature_a/dto"
	"hackathon-be/internal/app/feature_a/entity"

	"gorm.io/gorm"
)

type FeatureARepository interface {
	WithTx(tx *gorm.DB) FeatureARepository
	Create(ctx context.Context, e *entity.FeatureA) error
	FindByID(ctx context.Context, id uint) (*entity.FeatureA, error)
	List(ctx context.Context, q dto.ListFeatureAQuery) ([]entity.FeatureA, int64, error)
	Update(ctx context.Context, e *entity.FeatureA) error
	Delete(ctx context.Context, id uint) error
	// TODO: method khusus, contoh conditional update:
	// UpdateIfStatus(ctx, id, ownerID uint, fromStatus string, fields map[string]any) (rows int64, err error)
}
