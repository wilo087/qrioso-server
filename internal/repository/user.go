package repository

import (
	"github.com/google/uuid"
	"github.com/wilo0087/qrioso-server/internal/model"
	"github.com/wilo0087/qrioso-server/internal/model/dto"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUserByID(userID uuid.UUID) (*dto.UserResponse, error) {
	user := &model.User{}
	err := ur.db.Select("id, first_name, last_name, gender, birthdate, document_type, document, picture, role, created_at, updated_at").
		Preload("Emails").
		Preload("Companies").
		Preload("SocialNetworks").
		Where("id = ?", userID).
		First(user).Error

	if err != nil {
		return nil, err
	}

	userResponse := &dto.UserResponse{
		ID:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Emails:         user.Emails,
		Gender:         user.Gender,
		Birthdate:      user.Birthdate,
		DocumentType:   user.DocumentType,
		Document:       user.Document,
		Picture:        user.Picture,
		Role:           user.Role,
		Companies:      user.Companies,
		SocialNetworks: user.SocialNetworks,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}

	return userResponse, nil
}
