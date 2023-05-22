package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"

	"github.com/MSFT/internal/cfg"
)

func RecieveOrder(c *cfg.Config) (<-chan amqp091.Delivery, error) {
	q, err := Chann.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	err = Chann.QueueBind(
		q.Name,                // queue name
		"",                    // routing key
		c.Rabbitmq_queue_name, // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	msgs, err := Chann.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}
