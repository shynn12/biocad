package consumer

import (
	"context"
	"fmt"
	"os"

	"github.com/shynn12/biocad/internal/config"
	"github.com/shynn12/biocad/internal/item"
	"github.com/shynn12/biocad/pkg/logging"
	"github.com/shynn12/biocad/pkg/parser/tsv"
	"github.com/shynn12/biocad/pkg/pdfmaker"
	"github.com/streadway/amqp"
)

func CreateTSV(logger *logging.Logger, ch *amqp.Channel, s item.Service, cfg *config.Config) {
	msgs, err := ch.Consume(
		"TSVS",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Error(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			logger.Infof("Recieved message: %s", d.Body)
			f, err := os.Open(fmt.Sprintf("%s/%s", cfg.Tsvpath, string(d.Body)))
			if err != nil {
				logger.Errorf("cannot open a file due to error: %v", err)
			}
			items := tsv.Parse(f, logger)
			f.Close()
			for _, item := range items {
				sitem := []string{}
				s.CreateItem(context.Background(), item)
				sitem = append(sitem, item.Number, item.Mqtt, item.Invid, item.UnitGuid,
					item.MsgId, item.Text, item.Context, item.Class, item.Level, item.Area,
					item.Addr, item.Block, item.Type, item.Bit, item.InvertBit)
				err = pdfmaker.MakePDF(cfg.Headers, sitem, cfg)
				if err != nil {
					logger.Error(err)
				}
			}
		}
	}()

	logger.Info("Successfuly connected to RabbitMQ instance")

	logger.Info("Waiting for messages")
	<-forever

}
