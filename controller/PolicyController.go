package controller

import (
	"encoding/json"
	"log"
	"producer/database"
	"producer/model"
	"producer/rabbitmq"
	"producer/util"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InsertPolicy(c *fiber.Ctx) error {

	log.Default().Println("request new policy")

	var policy model.Policy

	if err := c.BodyParser(&policy); err != nil {
		return err
	}

	policy.CreatedAt = time.Now()

	//insert policy
	if result := database.DB.Create(&policy); result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	//convert policy data to JSON format for RabbitMQ message
	policyJSON, err := json.Marshal(policy)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to encode policy data !")
	}

	//publish message to RabbitMQ queue
	if err := rabbitmq.PublishMessage("policy_queue", string(policyJSON)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to publish message to RabbitMQ")
	}

	return util.SendResponse(c, "Success", policy, fiber.StatusCreated)
}

func UpdatePolicy(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Policy ID !")
	}

	policy := model.Policy{
		Id: uint(id),
	}

	if err := c.BodyParser(&policy); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request Body !")
	}

	if result := database.DB.Model(&policy).Updates(policy); result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed Update Data !")
	}

	return util.SendResponse(c, "Success", policy, fiber.StatusOK)
}

func GetPolicy(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Policy ID !")
	}

	var policy model.Policy

	result := database.DB.First(&policy, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "Policy By ID Notfound !")
		}
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	return util.SendResponse(c, "success", policy, fiber.StatusAccepted)
}

func DeletePolicy(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Policy ID !")
	}

	var policy model.Policy

	result := database.DB.First(&policy, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, "Policy By ID Notfound !")
		}
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	if result := database.DB.Delete(&policy); result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	return util.SendResponse(c, "success", id, fiber.StatusAccepted)
}
