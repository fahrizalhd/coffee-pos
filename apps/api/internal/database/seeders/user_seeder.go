package seeders

import (
	"coffee-pos-api/internal/modules/user"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB) error {
	var existingUser user.User

	err := db.
		Where("email = ?", "admin@coffeepos.dev").
		First(&existingUser).
		Error

	if err == nil {
		return nil
	}

	password, err := bcrypt.GenerateFromPassword(
		[]byte("admin123"),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	admin := user.User{
		Name:     "Administrator",
		Email:    "admin@coffeepos.dev",
		Password: string(password),
		Role:     "admin",
	}

	return db.Create(&admin).Error
}
