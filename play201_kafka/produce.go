package play201_kafka

import (
	"fmt"
	"github.com/IBM/sarama"
)

func Produce() {
	// 设置 Kafka 生产者配置属性
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	// 创建 Kafka 生产者
	producer, err := sarama.NewSyncProducer([]string{ADDR}, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	// 发送消息
	for i := 0; i < 100*1000; i++ {
		key := "key2"
		value := "value2"
		message := &sarama.ProducerMessage{
			Topic: TOPIC,
			Key:   sarama.StringEncoder(key),
			Value: sarama.StringEncoder(value),
		}
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	}
}
