package rabbitmq

import (
	"fmt"

	"github.com/shynn12/biocad/internal/config"
	"github.com/shynn12/biocad/pkg/logging"
	"github.com/streadway/amqp"
)

func RbmInit(cfg *config.Config, logger *logging.Logger) (conn *amqp.Connection, ch *amqp.Channel, err error) {
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		return nil, nil, fmt.Errorf("can`t connect to rbm due to error: %v", err)
	}

	ch, err = conn.Channel()
	if err != nil {
		return nil, nil, fmt.Errorf("can`t connect to rbm due to error: %v", err)
	}

	q, err := ch.QueueDeclare(
		"TSVS",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, nil, fmt.Errorf("cannot configure a queue due to error %v", err)
	}

	logger.Info(q)

	return conn, ch, nil
}
