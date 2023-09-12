package play032_rabbitmq_ttl

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Play() {
	// 连接到RabbitMQ服务器
	conn, err := amqp.Dial("amqp://guest:guest@192.168.31.8:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明延时交换机，需要使用插件提供的 x-delayed-message 类型
	err = ch.ExchangeDeclare("delayed-exchange", "x-delayed-message", true, false, false, false, nil)
	failOnError(err, "Failed to declare the exchange")

	// 声明一个队列
	queueName := "delayed-queue"
	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	// 绑定队列到延时交换机，并指定延时的时间（以毫秒为单位）
	err = ch.QueueBind(queueName, "", "delayed-exchange", false, nil)
	failOnError(err, "Failed to bind queue")

	// 发送延时消息
	message := "Hello, delayed world!"
	delayedTime := 5000 // 延时5秒（以毫秒为单位）

	headers := make(amqp.Table)
	headers["x-delay"] = delayedTime

	err = ch.Publish(
		"delayed-exchange", // 交换机名称
		"",                 // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
			Headers:     headers,
		})
	failOnError(err, "Failed to publish a delayed message")

	fmt.Printf("Sent: %s\n", message)

	// 接收延时消息
	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	fmt.Println("Waiting for the delayed message...")

	for msg := range msgs {
		fmt.Printf("Received: %s\n", msg.Body)
	}

}
