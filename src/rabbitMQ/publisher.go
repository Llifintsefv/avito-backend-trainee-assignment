package rabbitmq

import (
	"github.com/streadway/amqp"
)

type Publisher struct {
	channel *amqp.Channel
}
func NewPublisher(ch *amqp.Channel) (*Publisher, error) {
	return &Publisher{channel: ch}, nil
}

func (p *Publisher) PublishGenerateValue(reqJson []byte) error{
	
	err := p.channel.Publish(
		"GenerationQueue",
		"GenerateValue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        reqJson,
		},
	)
	if err != nil {
		return err
	}
	return nil

	
}