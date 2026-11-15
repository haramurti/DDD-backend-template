package service

import (
	"context"

	"hackathon-be/internal/app/feature_a/contract"
	"hackathon-be/internal/app/feature_a/dto"
	"hackathon-be/internal/realtime"
	"hackathon-be/pkg/response"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type service struct {
	repo contract.FeatureARepository
	db   *gorm.DB      // buat transaksi
	rdb  *redis.Client // boleh nil (redis mati -> skip cache)
	hub  *realtime.Hub
}

func New(repo contract.FeatureARepository, db *gorm.DB, rdb *redis.Client, hub *realtime.Hub) contract.FeatureAService {
	return &service{repo: repo, db: db, rdb: rdb, hub: hub}
}

// Aturan bisnis & hak akses ditaruh di SERVICE, bukan handler.

func (s *service) Create(ctx context.Context, actorID uint, req dto.CreateFeatureARequest) (*dto.FeatureAResponse, error) {
	// TODO: validasi aturan bisnis -> repo.Create -> invalidasi cache -> broadcast realtime
	return nil, response.ErrNotImplemented
}

func (s *service) GetByID(ctx context.Context, actorID, id uint) (*dto.FeatureAResponse, error) {
	// TODO: repo.FindByID -> cek akses (owner/role/membership) -> dto.FromFeatureA
	return nil, response.ErrNotImplemented
}

func (s *service) List(ctx context.Context, actorID uint, q dto.ListFeatureAQuery) ([]dto.FeatureAResponse, int64, error) {
	// TODO: cache-aside (data bersama aja) -> repo.List -> map ke dto
	return nil, 0, response.ErrNotImplemented
}

func (s *service) Update(ctx context.Context, actorID, id uint, req dto.UpdateFeatureARequest) (*dto.FeatureAResponse, error) {
	// TODO: cek owner + status secara ATOMIK (conditional update / lock), 409 kalau status udah berubah
	return nil, response.ErrNotImplemented
}

func (s *service) Delete(ctx context.Context, actorID, id uint) error {
	// TODO: cek akses -> delete (+ data turunan dalam transaksi) -> invalidasi cache -> broadcast
	return response.ErrNotImplemented
}
