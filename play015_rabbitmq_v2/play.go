package play015_rabbitmq_v2

import (
	"fmt"
	"github.com/streadway/amqp"
	"go-playground/util/util_mq"
	"log"
	"time"
)

const (
	RABBIT_URL   = "amqp://guest:guest@192.168.31.117:5672/" // RabbitMQ连接URL
	QUEUE_NAME   = "playground"
	PRIORITY_MAX = 10
)

func PlayAsProducerBlock() {
	SIZE_PRODUCER := 1

	mq := util_mq.RabbitMQManager{}
	defer mq.Release()

	if err := mq.Initialize(RABBIT_URL, SIZE_PRODUCER, false); err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	log.Println("PlayAsProducerBlock init ok")
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
					if err := mq.Initialize(RABBIT_URL, SIZE_PRODUCER, false); err != nil {
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
			reason, ok := <-mq.GetConn().NotifyClose(make(chan *amqp.Error))
			if ok {
				ch <- reason
			}
		}

		log.Println("NotifyClose end")
	}()

	time.Sleep(time.Second * 2)
	mq.DeclarePublishQueue(QUEUE_NAME, PRIORITY_MAX)
	for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 1)
		s := fmt.Sprintf("i := %8d", i)
		//log.Println(s)
		mq.Publish(QUEUE_NAME, []byte(s), uint8(4))
	}

	log.Println("waiting select")
	select {}
}

func PlayAsConsumerBlock(consumerCnt int) {
	mq := util_mq.RabbitMQManager{}
	defer mq.Release()

	if err := mq.Initialize(RABBIT_URL, consumerCnt, true); err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	log.Println("PlayAsConsumerBlock init ok")
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
					if err := mq.Initialize(RABBIT_URL, consumerCnt, true); err != nil {
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
			reason, ok := <-mq.GetConn().NotifyClose(make(chan *amqp.Error))
			if ok {
				ch <- reason
			}
		}

		log.Println("NotifyClose end")
	}()

	for i := 0; i < mq.GetSize(); i++ {
		log.Println("Consume ...")
		go mq.Consume(
			QUEUE_NAME,
			i,
			func(body []byte) bool {
				//log.Println("shouldNack body: ", string(body))
				return false
			})
	}

	// 等待程序退出
	log.Println("waiting select")
	select {}
}
