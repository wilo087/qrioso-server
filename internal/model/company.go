package model

import (
	"time"
)

type Company struct {
	ID               uint64           `gorm:"primaryKey;auto_increment"`
	Name             string           `gorm:"not null"`
	Address          string           `gorm:"default:null"`
	Phone            string           `gorm:"not null"`
	Email            string           `gorm:"not null"`
	Website          string           `gorm:"default:null"`
	Users            []User           `gorm:"many2many:users_on_companies"`
	UsersOnCompanies []UsersOnCompany `gorm:"foreignKey:CompanyID"`
	CreatedAt        time.Time        `gorm:"default:current_timestamp;column:created_at"`
	UpdatedAt        time.Time        `gorm:"autoUpdateTime;column:updated_at"`
}
