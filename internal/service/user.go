package service

import (
	"github.com/google/uuid"
	"github.com/wilo0087/qrioso-server/internal/model/dto"
	"github.com/wilo0087/qrioso-server/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{userRepository: ur}
}

func (u *UserService) GetUserByID(id uuid.UUID) (*dto.UserResponse, error) {
	user, err := u.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) CreateUser(user *dto.UserRequest) (*dto.UserResponse, error) {
	createdUser, err := u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
