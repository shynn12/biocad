package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/shynn12/biocad/internal/config"
	"github.com/shynn12/biocad/internal/item"
	"github.com/shynn12/biocad/internal/item/db"
	"github.com/shynn12/biocad/internal/rabbitMQ/consumer"
	"github.com/shynn12/biocad/internal/rabbitMQ/publisher"
	"github.com/shynn12/biocad/pkg/broker/rabbitmq"
	"github.com/shynn12/biocad/pkg/client/mongodb"
	"github.com/shynn12/biocad/pkg/logging"
)

func main() {

	logger := logging.GetLogger()
	logger.Info("create router")
	router := mux.NewRouter()

	cfg := config.GetConfig()

	logger.Info("connceting DB")

	client, err := mongodb.NewClient(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.Auth_db)
	if err != nil {
		logger.Fatal(err)
	}

	storage := db.NewStorage(client, cfg.MongoDB.Collection, logger)

	service := item.NewService(storage, logger)

	logger.Info("initing RabbitMQ")
	conn, ch, err := rabbitmq.RbmInit(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	defer conn.Close()

	defer ch.Close()
	path := fmt.Sprintf("%s/", cfg.Tsvpath)
	go publisher.Checker(path, logger, ch)
	go consumer.CreateTSV(logger, ch, service, cfg)
	logger.Info("register user handler")

	handler := item.NewHandler(logger, service, cfg)
	handler.Register(router)

	start(router, logger, cfg)
}

func start(router *mux.Router, logger *logging.Logger, cfg *config.Config) {
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("listen unix socket")

		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket %s", socketPath)
	} else {
		logger.Info("Listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port))
		logger.Infof("server is listening port %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
