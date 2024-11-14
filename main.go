package main

import (
	"log"
	"producer/database"
	"producer/rabbitmq"
	"producer/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	//initiate database
	database.Connect()

	//connect rabbit MQ
	err := rabbitmq.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ : %s", err)
		return
	}

	defer rabbitmq.CloseRabbitMQConnection()

	//goroutine start consume queue
	go func() {
		if err := rabbitmq.ConsumeMessage("policy_queue"); err != nil {
			log.Fatalf("Failed to consumer : %s", err)
		}
	}()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))

	//setup routes
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
