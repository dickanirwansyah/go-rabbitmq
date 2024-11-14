package controller

import (
	"encoding/json"
	"log"
	"producer/database"
	"producer/model"
	"producer/rabbitmq"
	"producer/util"
	"time"

	"github.com/gofiber/fiber/v2"
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
