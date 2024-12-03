package controller

import (
	"fmt"
	"producer/database"
	"producer/model"
	"producer/util"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateNewRoles(c *fiber.Ctx) error {

	var req model.RequestRole

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Request"})
	}

	if req.RolesName == "" || len(req.PermissionRequestList) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Role name or Permissions is empty !"})
	}

	tx := database.DB.Begin()

	roles := model.Roles{Name: req.RolesName}
	if err := tx.Create(&roles).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed create roles !"})
	}

	for _, permissionID := range req.PermissionRequestList {
		var permissions model.Permissions
		if err := tx.First(&permissions, permissionID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				tx.Rollback()
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": fmt.Sprintf("Permission ID with ID %d does not exists", permissionID),
				})
			}
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to check permissions existing",
			})
		}

		permissionRole := model.PermissionsRoles{
			RolesId:       roles.Id,
			PermissionsId: permissionID,
		}

		if err := tx.Create(&permissionRole).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to associate permissions",
			})
		}
	}

	//commit
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to commit transaction",
		})
	}

	return util.SendResponse(c, "Success", req, fiber.StatusAccepted)
}

func UpdateNewRoles(c *fiber.Ctx) error {

	var req model.RequestRole

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request",
		})
	}

	if req.RolesId == 0 || req.RolesName == "" || len(req.PermissionRequestList) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request, Roles ID, Roles Name, Permission Cannot be empty !",
		})
	}

	tx := database.DB.Begin()

	var existingRoles model.Roles
	if err := tx.First(&existingRoles, "id = ?", req.RolesId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			tx.Rollback()
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Roles does not exists !"})
		}
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed find roles !",
		})
	}

	existingRoles.Name = req.RolesName
	if err := tx.Updates(existingRoles).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed update roles",
		})
	}

	//delete existing data permission by roles
	if err := tx.Where("roles_id = ?", existingRoles.Id).Delete(&model.PermissionsRoles{}).Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed delete permission by roles !",
		})
	}

	//add new permissions
	for _, permissionID := range req.PermissionRequestList {
		var permissions model.Permissions
		//check if permission is exists
		if err := tx.First(&permissions, permissionID).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				tx.Rollback()
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": fmt.Sprintf("Permission ID %s does not exists", permissionID),
				})
			}
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to check permission exising",
			})
		}

		//add permission roles
		permissionRoles := model.PermissionsRoles{
			RolesId:       existingRoles.Id,
			PermissionsId: permissionID,
		}
		if err := tx.Create(&permissionRoles).Error; err != nil {
			tx.Rollback()
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to associate permissions ",
			})
		}
	}

	//commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to commit save permission roles",
		})
	}

	return util.SendResponse(c, "Success", req, fiber.StatusAccepted)
}
