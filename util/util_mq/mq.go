package util_mq

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"log"
	"time"
)

func failOnError(err error, msg string) { // optimize later
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type RabbitMQManager struct {
	url            string
	size           int
	conn           *amqp.Connection
	ch             []*amqp.Channel
	isConnected    bool
	reconnectTries int
}

func (r *RabbitMQManager) GetSize() int {
	return r.size
}

func (r *RabbitMQManager) GetConn() *amqp.Connection {
	return r.conn
}

func (r *RabbitMQManager) DeclarePublishQueue(queueName string, priorityMax int64) error {
	_, err := r.ch[0].QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		amqp.Table{"x-max-priority": uint8(priorityMax)}, // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return nil
}

func (r *RabbitMQManager) Publish(queueName string, body []byte, priority uint8) error {

	err := r.ch[0].Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Priority:    priority,
			Body:        body,
		})
	failOnError(err, "Publish")

	//log.Printf(" [x] Sent %s\n", body)

	return nil
}

func (r *RabbitMQManager) Consume(queueName string, channelId int, shouldNack func([]byte) bool) error {
	msgs, err := r.ch[channelId].Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	failOnError(err, "Failed to register a consumer")
	for d := range msgs {
		if shouldNack(d.Body) {
			d.Nack(false, true)
		} else {
			d.Ack(false)
		}
	}

	return nil
}

func (r *RabbitMQManager) Initialize(rabbitMQURL string, size int, isCons bool) error {
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
			return errors.Wrapf(err, "conn.Channel")
		}

		if isCons {
			//err = ch.Qos(1, 0, false)
			//if err != nil {
			//	return errors.Wrapf(err, "ch.Qos")
			//}
		}

		r.ch = append(r.ch, ch)
	}

	r.url = rabbitMQURL
	r.size = size
	r.conn = conn
	r.isConnected = true
	r.reconnectTries = 0

	log.Println("RabbitMQ connection initialized")

	ch := make(chan *amqp.Error)

	go func() {
		const timeout = 5 * time.Second
		timer := time.NewTimer(timeout)
		for {
			select {
			case d, ok := <-ch:
				if ok {
					log.Println("d: ", d)
					time.Sleep(time.Second * 3)
					if err := r.Initialize(r.url, r.size, isCons); err != nil {
						log.Fatalf("Failed to initialize RabbitMQ: %v", err)
					}
					log.Println("init ok")
				}
			case <-timer.C:
				//log.Println("timer.C: ", timer.C)
				timer.Reset(timeout)
			}
		}

		log.Println("select end")
	}()

	go func() {
		for {
			reason, ok := <-r.conn.NotifyClose(make(chan *amqp.Error))
			if ok {
				ch <- reason
			}
		}

		log.Println("NotifyClose end")
	}()

	return nil
}

/*
func (r *RabbitMQManager) reconnect(url string, size int) {
	r.isConnected = false
	for !r.isConnected {
		if err := r.initialize(url, size); err == nil {
			log.Printf("Failed to reconnect to RabbitMQ: %v", err)
			return
		}

		r.reconnectTries++

		reconnectInterval := time.Duration(1<<uint(r.reconnectTries)) * time.Second
		if reconnectInterval > 60*time.Second {
			reconnectInterval = 60 * time.Second
		}

		log.Printf("Retrying in %v...", reconnectInterval)
		time.Sleep(reconnectInterval)
	}
}
*/

func (r *RabbitMQManager) Release() {
	log.Println("release...")
	r.conn.Close()

	for i := 0; i < len(r.ch); i++ {
		r.ch[i].Close()
	}

	//defer rabbitMQ.conn.Close()
	//defer rabbitMQ.ch.Close()
}
