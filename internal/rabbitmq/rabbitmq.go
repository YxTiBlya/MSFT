package rabbitmq

import (
	"fmt"

	"github.com/MSFT/internal/cfg"
	amqp "github.com/rabbitmq/amqp091-go"
)

var Chann *amqp.Channel

func ConnToRabbitMQ(c *cfg.Config) error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%v:%d/", c.Rabbitmq_host, c.Rabbitmq_port))
	if err != nil {
		return err
	}

	Chann, err = conn.Channel()
	if err != nil {
		return err
	}

	if err = Chann.ExchangeDeclare(
		c.Rabbitmq_queue_name,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return err
	}

	return nil
}
