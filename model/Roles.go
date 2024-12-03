package model

type Roles struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"varchar(150)"`
}
