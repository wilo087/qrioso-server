package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"size:255"`
	Address   string    `gorm:"size:255"`
	Website   string    `gorm:"size:255"`
	Status    CompanyStatus
	Users     []*User          `gorm:"many2many:user_on_companies"`
	Configs   []*CompanyConfig `gorm:"foreignKey:CompanyID"`
	CreatedAt time.Time        `gorm:"index;not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt *time.Time       `gorm:"index" json:"updated_at"`
	DeletedAt gorm.DeletedAt   `gorm:"index"`
}

type UserOnCompany struct {
	CompanyID uuid.UUID           `gorm:"type:uuid;primaryKey" json:"company_id"`
	UserID    uuid.UUID           `gorm:"type:uuid;primaryKey" json:"user_id"`
	IsAdmin   bool                `gorm:"default:false" json:"is_admin"`
	Status    UserOnCompanyStatus `gorm:"default:active" json:"status"`
	CreatedAt time.Time           `gorm:"index;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time          `gorm:"index" json:"updated_at"`
	DeletedAt gorm.DeletedAt      `gorm:"index"`
}

type CompanyConfig struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"`
	CompanyID  uuid.UUID `gorm:"type:uuid"`
	Config     string
	ConfigType string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type CompanyStatus string

const (
	CompanyStatusVerified    CompanyStatus = "verified"
	CompanyStatusActive      CompanyStatus = "active"
	CompanyStatusUnpublished CompanyStatus = "unpublished"
	CompanyStatusDeleted     CompanyStatus = "deleted"
)

type UserOnCompanyStatus string

const (
	UserOnCompanyStatusDisable UserOnCompanyStatus = "disabled"
	UserOnCompanyStatusActive  UserOnCompanyStatus = "active"
	UserOnCompanyStatusDeleted UserOnCompanyStatus = "deleted"
)

// func (UserOnCompany) BeforeCreate(db *gorm.DB) error {
// 	// ...
// }
