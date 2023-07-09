package model

import (
	"time"

	"github.com/google/uuid"
)

type SocialNetwork struct {
	UUID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"not null"`
	URL       string    `gorm:"not null"`
	User      User      `gorm:"foreignKey:UserID"`
	UserID    uint64    `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"default:current_timestamp;column:created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:updated_at"`
}
