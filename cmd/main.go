package main

import (
	"github.com/anilozgok/gp-prototype-consumer/internal/rabbit"
	"go.uber.org/zap"
	"log"
)

func main() {

	rabbitClient, err := rabbit.New()
	if err != nil {
		log.Fatal("failed to create rabbit client", zap.Error(err))
	}
	defer rabbitClient.CloseConnection()

	if err = rabbitClient.OpenChannel(); err != nil {
		log.Fatal("failed to open channel", zap.Error(err))
	}
	defer rabbitClient.CloseChannel()

	msgs, err := rabbitClient.Ch.Consume(
		"test",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("failed to consume message", zap.Error(err))
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received message: %s", d.Body)
		}
	}()
	<-forever
}
