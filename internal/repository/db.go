package repository

import (
	"github.com/wilo0087/qrioso-server/cmd/database"
	"github.com/wilo0087/qrioso-server/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=qrioso port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&model.User{},
		&model.Company{},
		&model.UserOnCompany{},
		&model.UserEmail{},
		&model.SocialNetwork{},
	)

	// Many to many relationship with custom fields
	db.SetupJoinTable(&model.User{}, "Companies", &model.UserOnCompany{})

	// Create default company and user
	database.DefaultCompanyAndUser(db)

	return db
}
