package database

import (
	featureAEntity "hackathon-be/internal/app/feature_a/entity"
	featureBEntity "hackathon-be/internal/app/feature_b/entity"
	userEntity "hackathon-be/internal/app/user/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&userEntity.User{},
		&featureAEntity.FeatureA{},
		&featureBEntity.FeatureB{},
		// TODO: daftarin entity baru di sini
	)
}
