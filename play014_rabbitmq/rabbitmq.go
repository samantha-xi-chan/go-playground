package play014_rabbitmq

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQManager struct {
	url            string
	size           int
	conn           *amqp.Connection
	ch             []*amqp.Channel
	isConnected    bool
	reconnectTries int
}

func (r *RabbitMQManager) initialize(rabbitMQURL string, size int) error {
	conn, err := amqp.DialConfig(
		rabbitMQURL,
		amqp.Config{
			Heartbeat: 1 * time.Second,
		})
	if err != nil {
		return err
	}

	r.ch = nil
	for i := 0; i < size; i++ {
		ch, err := conn.Channel()
		if err != nil {
			return err
		}
		r.ch = append(r.ch, ch)
	}

	r.url = rabbitMQURL
	r.size = size
	r.conn = conn
	r.isConnected = true
	r.reconnectTries = 0

	log.Println("RabbitMQ connection initialized")
	return nil
}

func (r *RabbitMQManager) release() {
	log.Println("release...")
	r.conn.Close()

	for i := 0; i < len(r.ch); i++ {
		r.ch[i].Close()
	}

	//defer rabbitMQ.conn.Close()
	//defer rabbitMQ.ch.Close()
}
