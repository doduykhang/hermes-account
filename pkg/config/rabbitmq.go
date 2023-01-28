package config

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func getRabbitMQConnString(rabbitMQ RabbitMQ) string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		rabbitMQ.User,
		rabbitMQ.Password,
		rabbitMQ.Host,
		rabbitMQ.Port,
	)
}

func NewRabbitMq(config *Config) *amqp.Connection {
	connString := getRabbitMQConnString(config.RabbitMQ)
	conn, err := amqp.Dial(connString)
	if err != nil {
		log.Panic(err)
	}
	return conn
}
