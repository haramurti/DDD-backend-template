package database

import "gorm.io/gorm"

// Seed bikin data demo (akun per role, data awal) buat demo & pitching.
func Seed(db *gorm.DB) error {
	// TODO: buat akun admin/pengelola + beberapa data contoh
	_ = db
	return nil
}
