package repository

import (
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/wilo0087/qrioso-server/internal/model"
	"github.com/wilo0087/qrioso-server/internal/model/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUserByID(userID uuid.UUID) (*dto.UserResponse, error) {
	res := dto.UserResponse{}
	user := &model.User{}
	err := ur.db.Select("id, first_name, last_name, gender, birthdate, document_type, document, picture, role, created_at, updated_at").
		Preload(clause.Associations).
		Where("id = ?", userID).
		First(user).Error

	if err != nil {
		return nil, err
	}

	copier.Copy(&res, &user)
	return &res, nil
}

func (ur *UserRepository) CreateUser(u *dto.UserRequest) (*dto.UserResponse, error) {
	user := model.User{}
	copier.Copy(&user, &u)

	err := ur.db.Create(&user).Error

	if err != nil {
		return nil, err
	}

	res := dto.UserResponse{}
	copier.Copy(&res, &user)

	return &res, nil
}
