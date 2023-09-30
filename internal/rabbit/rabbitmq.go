package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
)

type RabbitClient struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func New() (*RabbitClient, error) {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		"guest",
		"guest",
		"localhost",
		5672,
	)

	conn, err := amqp.Dial(connStr)
	if err != nil {
		return nil, err
	}

	return &RabbitClient{Conn: conn}, nil
}

func (r *RabbitClient) CloseConnection() {
	r.Conn.Close()
}

func (r *RabbitClient) OpenChannel() error {
	ch, err := r.Conn.Channel()
	if err != nil {
		log.Fatal("failed to open a channel", zap.Error(err))
		return err
	}
	r.Ch = ch

	return nil
}

func (r *RabbitClient) CloseChannel() {
	r.Ch.Close()
}

func (r *RabbitClient) DeclareQueue(name string) error {
	_, err := r.Ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RabbitClient) PublishMessage(name string, message string) error {
	err := r.Ch.Publish(
		"",
		name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		return err
	}

	return nil
}
