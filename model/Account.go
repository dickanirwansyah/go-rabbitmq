package model

import "time"

type Account struct {
	Id          uint      `gorm:"primaryKey"`
	FullName    string    `gorm:"type:varchar(150)"`
	Email       string    `gorm:"type:varchar(150)"`
	PhoneNumber string    `gorm:"type:varchar(20)"`
	RoleCode    uint      `gorm:"type:int"`
	Roles       Roles     `gorm:"foreignKey:RoleCode;references:Id"`
	CreatedAt   time.Time `gorm:"type:date"`
	UpdatedAt   time.Time `gorm:"type:date"`
}
