package config

import "os"

type Config struct {
	Port        string
	DatabaseURL string
	RedisAddr   string
	JWTSecret   string
	CORSOrigin  string
	// TODO: tambah config lain sesuai kebutuhan (MinIO, Google OAuth client ID, dll)
}

func Load() Config {
	return Config{
		Port:        getenv("PORT", "8080"),
		DatabaseURL: getenv("DATABASE_URL", "host=localhost user=postgres password=postgres dbname=app port=5432 sslmode=disable TimeZone=Asia/Jakarta"),
		RedisAddr:   getenv("REDIS_ADDR", "localhost:6379"),
		JWTSecret:   getenv("JWT_SECRET", "dev-secret-ganti"),
		CORSOrigin:  getenv("CORS_ORIGIN", "http://localhost:3000"),
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
