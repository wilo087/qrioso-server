package service

import (
	"github.com/wilo0087/qrioso-server/internal/model"
	"github.com/wilo0087/qrioso-server/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{userRepository: ur}
}

func (u *UserService) GetUserByID(id string) (*model.User, error) {
	user, err := u.userRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
