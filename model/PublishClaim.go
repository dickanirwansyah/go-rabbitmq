package model

import (
	"time"
)

type PublishClaim struct {
	Id             uint      `gorm:"primaryKey"`
	PolicyNumber   string    `gorm:"type:varchar(150)"`
	ClaimInsurance string    `gorm:"type:varchar(100)"`
	ClaimDate      time.Time `gorm:"type:date"`
}
