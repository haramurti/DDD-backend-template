package storage

import (
	"context"
	"io"
	"time"
)

// Storage object storage (MinIO/S3). Hapus kalau PRD gak butuh upload (IdeaBox, Expense Tracker).
type Storage interface {
	Put(ctx context.Context, key string, r io.Reader, size int64, contentType string) error
	PresignGet(ctx context.Context, key string, ttl time.Duration) (string, error)
	Delete(ctx context.Context, key string) error
}

// TODO: implementasi pakai github.com/minio/minio-go/v7
// - bucket private, DB simpan object key (bukan URL)
// - PresignGet dipanggil SETELAH cek akses
