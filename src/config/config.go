package config

import "fmt"

type RabbitMQConfig struct {
	Host      string
	Port      int
	User      string
	Pass      string
	QueueName string
}

func GetRabbitConfig() RabbitMQConfig {
	return RabbitMQConfig{Host: "localhost", Port: 5672, User: "guest", Pass: "guest", QueueName: "random_values"}
}

func GetRabbutMQConnectionString(config RabbitMQConfig) string {
	cfg := GetRabbitConfig()

	return fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.User, cfg.Pass, cfg.Host, cfg.Port)
}