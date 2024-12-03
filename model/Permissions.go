package model

type Permissions struct {
	Id                  uint   `gorm:"primaryKey"`
	PermissionName      string `gorm:"type:varchar(150)"`
	PermissionLevel     int    `gorm:"type:int"`
	PermissionParentId  int    `gorm:"type:int"`
	PermissionEndpoint  string `gorm:"type:varchar(150)"`
	PermissionGlyphicon string `gorm:"type:varchar(150)"`
	PermissionActive    string `gorm:"type:varchar(20)"`
}
