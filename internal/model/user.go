package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	BaseUUID
	FirstName      string         `gorm:"column:first_name"`
	LastName       string         `gorm:"column:last_name"`
	Email          string         `gorm:"unique" json:"email"`
	Picture        *string        `gorm:"default:null" json:"picture"`
	Role           UserRole       `gorm:"default:public" json:"role"`
	Password       sql.NullString `gorm:"default:null" json:"password"`
	Companies      []Company      `gorm:"many2many:user_on_companies" json:"companies"`
	SocialNetworks []SocialNetwork
	BaseModelSoftDelete
}

type UserEmail struct {
	UUID      uint      `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	Email     string    `gorm:"size:255;unique"`
	Main      bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleRegisted UserRole = "registed"
	UserRolePublic   UserRole = "public"
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
