package model

type PermissionsRoles struct {
	Id            uint        `gorm:"primaryKey"`
	RolesId       uint        `gorm:"type:int"`
	Roles         Roles       `gorm:"foreignKey:RolesId;references:Id"`
	PermissionsId uint        `gorm:"type:int"`
	Permissions   Permissions `gorm:"foreignKey:PermissionsId;references:Id"`
}
