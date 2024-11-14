package util

import "github.com/gofiber/fiber/v2"

type GenericResponse struct {
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code,omitempty"`
}

func SendResponse(c *fiber.Ctx, message string, data interface{}, statusCode int) error {
	return c.Status(statusCode).JSON(GenericResponse{
		Message:    message,
		Data:       data,
		StatusCode: statusCode,
	})
}
