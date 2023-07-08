package repository

import (
	"github.com/wilo0087/qrioso-server/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := ur.db.First(&user, "uuid = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
