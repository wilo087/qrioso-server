package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID               uint64         `gorm:"primary_key;index;not null" json:"id"`
	FirstName        string         `gorm:"column:first_name"`
	LastName         string         `gorm:"column:last_name"`
	Email            string         `gorm:"unique" json:"email"`
	Picture          *string        `gorm:"default:null" json:"picture"`
	Role             Role           `gorm:"default:public" json:"role"`
	Password         sql.NullString `gorm:"default:null" json:"password"`
	Companies        []Company      `gorm:"many2many:users_on_companies" json:"companies"`
	SocialNetworks   []SocialNetwork
	UsersOnCompanies []UsersOnCompany `gorm:"foreignKey:UserID"`
	CreatedAt        time.Time        `gorm:"default:current_timestamp;column:created_at"`
	UpdatedAt        time.Time        `gorm:"autoUpdateTime;column:updated_at"`
}

type UsersOnCompany struct {
	Company   Company `gorm:"foreignKey:CompanyID"`
	CompanyID uint64  `gorm:"column:company_id"`
	User      User    `gorm:"foreignKey:UserID"`
	UserID    uint64  `gorm:"column:user_id"`
}

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleRegisted Role = "registed"
	RolePublic   Role = "public"
)

// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	record := &User{}

// 	if u.Email != "" && db.Debug().Where("email = ?", u.Email).First(record).Error == nil {
// 		return errors.New("email already exists")
// 	}

// 	if u.Password != "" {
// 		if pw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 11); err == nil {
// 			tx.Statement.SetColumn("password", pw)
// 		}
// 	}

// 	return
// }
