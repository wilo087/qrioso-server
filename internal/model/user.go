package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey"`
	FirstName      string    `gorm:"column:first_name"`
	LastName       string    `gorm:"column:last_name"`
	Emails         []UserEmail
	Gender         string         `gorm:"size:1" json:"gender"`
	Birthdate      *time.Time     `json:"birthdate"`
	DocumentType   *UserDocType   `gorm:"size:1" validate:"omitempty" json:"document_type"`
	Document       *string        `gorm:"size:16" validate:"omitempty" json:"document"`
	Picture        *string        `gorm:"default:null" json:"picture"`
	Role           UserRole       `gorm:"default:public" json:"role"`
	Password       sql.NullString `gorm:"default:null" json:"password"`
	Companies      []*Company     `gorm:"many2many:user_on_companies" json:"companies"`
	SocialNetworks []*SocialNetwork
	CreatedAt      time.Time      `gorm:"default:current_timestamp;column:created_at"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

type UserEmail struct {
	Email     string `gorm:"size:255;primarykey"`
	Main      bool
	UserID    uuid.UUID `gorm:"type:uuid"`
	User      User
	CreatedAt time.Time      `gorm:"default:current_timestamp;column:created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

type UserRole string

const (
	UserRoleAdmin    UserRole = "admin"
	UserRoleRegisted UserRole = "registed"
	UserRolePublic   UserRole = "public"
)

type UserDocType string

const (
	UserDocTypeC        UserDocType = "C"
	UserDocTypePassport UserDocType = "P"
	UserDocTypeSocial   UserDocType = "S"
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
