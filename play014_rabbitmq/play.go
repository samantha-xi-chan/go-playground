package play014_rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"time"
)

const (
	RABBIT_URL = "amqp://guest:guest@192.168.31.8:5672/" // RabbitMQ连接URL
	SIZE       = 10
)

func Play() {
	mq := RabbitMQManager{}
	defer mq.release()

	if err := mq.initialize(RABBIT_URL, SIZE); err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	log.Println("init ok")
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
					if err := mq.initialize(RABBIT_URL, SIZE); err != nil {
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
			reason, ok := <-mq.conn.NotifyClose(make(chan *amqp.Error))
			if ok {
				ch <- reason
			}
		}

		log.Println("NotifyClose end")
	}()

	//rabbitMQ.ensureConnected()

	// 在这里可以使用rabbitMQ.ch来进行消息的发送和接收等操作
	// 例如：rabbitMQ.ch.Publish() 或 rabbitMQ.ch.Consume()

	// 等待程序退出

	log.Println("waiting select")
	//select {}
}
