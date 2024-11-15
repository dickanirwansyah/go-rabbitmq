package model

import (
	"time"
)

type PublishClaim struct {
	Id             uint   `gorm:"primaryKey"`
	PolicyNumber   string `gorm:"type:varhar(150)"`
	ClaimInsurance string `gorm:"type:varchar(100)"`
	ClaimDate      time.Time
}
