package rabbitmq

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/MSFT/internal/cfg"
	"github.com/MSFT/pkg/services/customer"
)

func RecieveOrder(c *cfg.Config) {
	q, err := Chann.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalln("rabbitmq: error to declare queue:\n", err.Error())
	}

	err = Chann.QueueBind(
		q.Name,                // queue name
		"",                    // routing key
		c.Rabbitmq_queue_name, // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatalln("rabbitmq: error to bind queue:\n", err.Error())
	}

	msgs, err := Chann.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalln("rabbitmq: error to consume init:\n", err.Error())
	}

	var forever chan struct{}
	log.Infoln("starting listening order queue at rabbitmq")
	go func() {
		for d := range msgs {
			log.Infof("recieved message: %s", d.Body)
			data := customer.CreateOrderRequest{}
			json.Unmarshal(d.Body, &data)
			//log.Println(data)
		}
	}()

	<-forever
}
