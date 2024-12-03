package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func InitRabbitMQ() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return  nil, err
	}

	err = ch.ExchangeDeclare(
		"GenerationQueue",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return  nil, err
	}

	_, err = ch.QueueDeclare(
		"GenerateValue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return  nil, err
	}

	err = ch.QueueBind(
		"GenerateValue",
		"GenerateValue",
		"GenerationQueue",
		false,
		nil,
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return  nil, err
	}

	log.Println("RabbitMQ connection and channel initialized successfully")
	return ch, nil
}