package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/wilo0087/qrioso-server/internal/model"
)

type CreateUserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Picture   string `json:"picture"`
	Role      string `json:"role"`
	Password  string `json:"password"`
}

type UserResponse struct {
	ID             uuid.UUID              `json:"id"`
	FirstName      string                 `json:"first_name"`
	LastName       string                 `json:"last_name"`
	Emails         []UserEmailResponse    `json:"emails"`
	Gender         string                 `json:"gender"`
	Birthdate      *time.Time             `json:"birthdate"`
	DocumentType   *model.UserDocType     `json:"document_type"`
	Document       *string                `json:"document"`
	Picture        *string                `json:"picture"`
	Role           model.UserRole         `json:"role"`
	Companies      []*model.Company       `json:"companies"`
	SocialNetworks []*model.SocialNetwork `json:"social_networks"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
}

type UserEmailResponse struct {
	Email     string    `json:"email"`
	Main      bool      `json:"main"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
