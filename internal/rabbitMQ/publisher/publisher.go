package publisher

import (
	"os"
	"time"

	"github.com/shynn12/biocad/pkg/logging"
	"github.com/shynn12/biocad/pkg/utilites"
	"github.com/streadway/amqp"
)

var checked []string

func Checker(dir string, logger *logging.Logger, ch *amqp.Channel) {
	interval := time.Second
	ticker := time.NewTicker(interval)
	for range ticker.C {
		logger.Infof("checking...")
		d, err := os.Open(dir)
		if err != nil {
			logger.Errorf("can`t open directory due to error: %v", err)
		}

		fileNames, err := d.Readdirnames(-1)
		if err != nil {
			logger.Error(err)
		}
		d.Close()
		for _, i := range fileNames {
			if !utilites.IsInSlice(i, checked) {
				logger.Infof("Found a new element %s", i)
				checked = append(checked, i)
				err = ch.Publish(
					"",
					"TSVS",
					false,
					false,
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(i),
					},
				)
				if err != nil {
					logger.Error(err)
				}
			}
		}
	}
}
