package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	BaseUUID
	Name    string `gorm:"size:255"`
	Address string `gorm:"size:255"`
	Website string `gorm:"size:255"`
	Status  CompanyStatus
	Users   []User          `gorm:"many2many:user_on_companies"`
	Configs []CompanyConfig `gorm:"foreignKey:CompanyID"`
	BaseModelSoftDelete
}

type UserOnCompany struct {
	CompanyID uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"primaryKey"`
	IsAdmin   bool      `gorm:"default:false"`
	Status    UserOnCompanyStatus
	BaseModelSoftDelete
}

type CompanyConfig struct {
	UUID       uint      `gorm:"primaryKey"`
	CompanyID  uuid.UUID `gorm:"type:uuid"`
	Config     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	ConfigType string
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
	UserOnCompanyStatusDisable UserOnCompanyStatus = "disable"
	UserOnCompanyStatusActive  UserOnCompanyStatus = "active"
	UserOnCompanyStatusDeleted UserOnCompanyStatus = "deleted"
)
