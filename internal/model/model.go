package model

import (
	"time"
)

// BaseID defines the common columns that all db structs should hold.
type BaseID struct {
	ID int `gorm:"primary_key;index;not null"`
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
	DeletedAt *time.Time `sql:"index" json:"deletedAt"`
}
