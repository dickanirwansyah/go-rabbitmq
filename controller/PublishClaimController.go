package controller

import (
	"encoding/json"
	"log"
	"producer/model"
	"producer/rabbitmq"
	"producer/util"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SendPublishClaim(c *fiber.Ctx) error {

	log.Default().Println("send publish claim")

	var publishClaim model.PublishClaim

	if err := c.BodyParser(&publishClaim); err != nil {
		return fiber.NewError(fiber.StatusNotAcceptable, "Payload Claim Insurance is not valid !")
	}

	publishClaim.ClaimDate = time.Now()

	publishClaimJSON, err := json.Marshal(publishClaim)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error Processing Marshal JSON publish Claim Insurance !")
	}

	if err := rabbitmq.PublishMessage("publish_claim", string(publishClaimJSON)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed publish messaging to topic publish_claim !")
	}

	return util.SendResponse(c, "success", publishClaim, fiber.StatusOK)
}
