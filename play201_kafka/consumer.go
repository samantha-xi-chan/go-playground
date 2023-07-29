package play201_kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"os"
	"os/signal"
)

func Consume() {
	// 配置Kafka消费者
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// 创建Kafka消费者
	consumer, err := sarama.NewConsumer([]string{ADDR}, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			panic(err)
		}
	}()

	// 订阅Kafka主题
	partitionConsumer, err := consumer.ConsumePartition(TOPIC, 0, sarama.OffsetOldest) // sarama.OffsetNewest
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			panic(err)
		}
	}()

	// 等待消费者接收到信号
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// 从Kafka主题中读取消息并处理它们
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("接收到消息: topic=%s, partition=%d, offset=%d, key=%s, value=%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		case <-signals:
			return
		}
	}
}
