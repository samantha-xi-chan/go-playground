package play015_rabbitmq_v2

import (
	"fmt"
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

	// 定义操作对象
	mq := util_mq.RabbitMQManager{}
	defer mq.Release()

	if err := mq.Initialize(RABBIT_URL, SIZE_PRODUCER, false); err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	log.Println("PlayAsProducerBlock init ok")
	time.Sleep(time.Second * 2)

	// 发送消息
	mq.DeclarePublishQueue(QUEUE_NAME, PRIORITY_MAX)
	for i := 0; ; i++ {
		time.Sleep(time.Millisecond * 1)
		s := fmt.Sprintf("i := %8d", i)
		//log.Println(s)
		mq.Publish(QUEUE_NAME, []byte(s), uint8(4))
	}

	// 阻塞等待
	log.Println("waiting select")
	select {}
}

func PlayAsConsumerBlock(consumerCnt int) {

	// 定义操作对象
	mq := util_mq.RabbitMQManager{}
	defer mq.Release()

	if err := mq.Initialize(RABBIT_URL, consumerCnt, true); err != nil {
		log.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}
	log.Println("PlayAsConsumerBlock init ok")

	// 监听消息
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

	// 阻塞等待
	log.Println("waiting select")
	select {}
}
