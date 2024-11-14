package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"producer/model"

	"github.com/streadway/amqp"
)

var Channel *amqp.Channel
var Connection *amqp.Connection

func Connect() error {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ : %s ", err)
		return err
	}

	Connection = conn

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel : %s", err)
		return err
	}

	Channel = ch
	return nil
}

func DeclareQueue(queueName string) (amqp.Queue, error) {

	queue, err := Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to declare a queue : %s", err)
		return amqp.Queue{}, err
	}

	return queue, nil
}

func PublishMessage(queueName string, message string) error {

	queue, err := DeclareQueue(queueName)
	if err != nil {
		return err
	}

	err = Channel.Publish(
		"",         //exchange
		queue.Name, //routing key (queue name)
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		log.Fatalf("Failed to publish a message : %s", err)
		return err
	}

	log.Printf("Sent message : %s", message)
	return nil
}

func ConsumeMessage(queueName string) error {

	_, err := DeclareQueue(queueName)
	if err != nil {
		return err
	}

	//start receiving messages from queue
	messagesFromTopic, err := Channel.Consume(
		queueName, //queue name
		"",
		true, //auto acknowledge message after receiving
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("Failed to register consumer : %v", err)
	}

	forever := make(chan bool)

	//create goroutine
	go func() {
		for d := range messagesFromTopic {
			//parsing json to object entity Polis
			var policy model.Policy
			if err := json.Unmarshal(d.Body, &policy); err != nil {
				log.Printf("Error unmarshal/decoding JSON : %s", err)
				continue
			}
			//process data Polis messages
			log.Printf("Receive policy : %+v", policy)
		}
	}()

	log.Printf("Waiting for message on queue : %s", queueName)
	<-forever // Block until next an exit signal is received
	return nil
}

func CloseRabbitMQConnection() {
	if Channel != nil {
		Channel.Close()
	}
	if Connection != nil {
		Connection.Close()
	}
}
