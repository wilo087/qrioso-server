package database

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/google/uuid"
	"github.com/wilo0087/qrioso-server/internal/model"
)

func SeedCompanyAndUsers(db *gorm.DB) error {
	doc := "224-0011982-7"
	docType := model.UserDocTypeC
	birthdate := time.Date(1938, time.February, 3, 0, 0, 0, 0, time.UTC)

	users := []model.User{
		{
			ID:        uuid.New(),
			FirstName: "Adrian",
			LastName:  "Mendez",
			Gender:    "F",
			Emails: []model.UserEmail{
				{
					Email: "mendezadrian149@gmail.com",
					Main:  true,
				},
				{
					Email: "tucochita_rica@hotmail.com",
					Main:  false,
				},
				{
					Email: "elma_mape_neduro@hotmail.com",
					Main:  false,
				},
			},
			Companies: []*model.Company{
				{
					ID:      uuid.New(),
					Name:    "Qrioso",
					Address: "Address 2",
					Website: "qrioso.com",
					Status:  model.CompanyStatusActive,
				},
			},
			Birthdate:    &birthdate,
			DocumentType: &docType,
			Document:     &doc,
			Picture:      nil,
			Role:         model.UserRoleAdmin,
		},
	}

	userOnCompany := model.UserOnCompany{}

	return db.Transaction(func(tx *gorm.DB) error {
		db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&users)

		for _, user := range users {
			userOnCompany.CompanyID = user.Companies[0].ID
			userOnCompany.UserID = user.ID
			userOnCompany.IsAdmin = true
		}

		db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&userOnCompany)

		return nil
	})
}
