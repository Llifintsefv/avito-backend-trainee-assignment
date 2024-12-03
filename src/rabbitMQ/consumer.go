package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"pro-backend-trainee-assignment/src/models"
	"pro-backend-trainee-assignment/src/service"

	"github.com/streadway/amqp"
)

type Consumer struct {
	channel *amqp.Channel
	service service.Service
}

func NewConsumer(ch *amqp.Channel,service service.Service) (*Consumer,error) {
	 return &Consumer{
        channel: ch,
        service: service,
    }, nil
}

func (c *Consumer) ConsumeGeneratedValue() {
	fmt.Println("Start consuming GenerateValue messages")
	msgs, err := c.channel.Consume(
		"GenerateValue",
		"",
		false, 
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to register a consumer %w", err)
	}

	for d := range msgs {
		fmt.Println("Received GenerateValue message")
		var req models.GenRequest
		err := json.Unmarshal(d.Body, &req)
		if err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			d.Nack(false, false)
			continue
		}

		go func(delivery amqp.Delivery) {
			c.service.GenerateNumber(req)		
		}(d)
	}

	log.Println("Stopped consuming messages")
}