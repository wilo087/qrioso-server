package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseID defines the common columns that all db structs should hold.
type BaseUUID struct {
	UUID uuid.UUID `gorm:"type:uuid;primaryKey"`
}

// BaseModel defines the common columns for all db models
type BaseModel struct {
	CreatedAt time.Time  `gorm:"index;not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"index" json:"updatedAt"`
}

// BaseModelSoftDelete defines the common columns that all db structs should.
// This struct also defines the fields for GORM triggers to detect the entity should soft delete
type BaseModelSoftDelete struct {
	BaseModel
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
