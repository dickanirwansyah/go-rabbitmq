package model

import "time"

type Policy struct {
	Id               uint      `gorm:"primaryKey"`
	PolicyNumber     string    `gorm:"type:varchar(150)"`
	PolicyHolder     string    `gorm:"type:varchar(150)`
	Insured          string    `gorm:"type:varchar(200)"`
	Beneficiary      string    `gorm:"type:varchar(200)"`
	InsuranceCarrier string    `gorm:"type:varchar(200)"`
	Underwriter      string    `gorm:"type:varchar(200)"`
	CreatedAt        time.Time `gorm:"type:date"`
	UpdatedAt        time.Time `gorm:"type:date"`
}
