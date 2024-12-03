package model

type RequestRole struct {
	RolesId               uint   `json:"RolesID"`
	RolesName             string `json:"RolesName" binding:"required"`
	PermissionRequestList []uint `json:"PermissionRequestList" binding:"required"`
}
