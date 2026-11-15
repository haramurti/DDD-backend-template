package repository

import (
	"context"

	"hackathon-be/internal/app/feature_b/contract"
	"hackathon-be/internal/app/feature_b/dto"
	"hackathon-be/internal/app/feature_b/entity"
	"hackathon-be/pkg/response"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) contract.FeatureBRepository {
	return &repository{db: db}
}

// WithTx dipakai service saat butuh transaksi: s.db.Transaction(func(tx) { repo.WithTx(tx)... })
func (r *repository) WithTx(tx *gorm.DB) contract.FeatureBRepository {
	return &repository{db: tx}
}

func (r *repository) Create(ctx context.Context, e *entity.FeatureB) error {
	// TODO: return r.db.WithContext(ctx).Create(e).Error
	return response.ErrNotImplemented
}

func (r *repository) FindByID(ctx context.Context, id uint) (*entity.FeatureB, error) {
	// TODO: var e entity.FeatureB; err := r.db.WithContext(ctx).First(&e, id).Error
	return nil, response.ErrNotImplemented
}

func (r *repository) List(ctx context.Context, q dto.ListFeatureBQuery) ([]entity.FeatureB, int64, error) {
	// TODO: filter (Where ... ?), Count, Order (whitelist!), Offset/Limit
	return nil, 0, response.ErrNotImplemented
}

func (r *repository) Update(ctx context.Context, e *entity.FeatureB) error {
	// TODO: pakai Updates(map[string]any{...}) biar field nol tetap ke-update
	return response.ErrNotImplemented
}

func (r *repository) Delete(ctx context.Context, id uint) error {
	// TODO: return r.db.WithContext(ctx).Delete(&entity.FeatureB{}, id).Error
	return response.ErrNotImplemented
}
