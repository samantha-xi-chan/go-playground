package play033_rmq_ttl

import (
	"fmt"
	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/jinzhu/configor"
	"go-playground/play033_rmq_ttl/rmq_util"
	"time"
)

func Play() {
	log.SetHandler(cli.Default)
	log.SetLevel(log.DebugLevel)

	var config rmq_util.Config

	err := configor.Load(&config, "config/mq.json")
	if err != nil {
		log.Fatalf("run: failed to init config: %v", err)
	}

	rmq, err := rmq_util.InitRabbitMQ(config.AMQP)
	if err != nil {
		log.Fatalf("run: failed to init rabbitmq: %v", err)
	}
	defer rmq.Shutdown()

	for i := 10; i > 0; i-- {
		msg := fmt.Sprintf(" i= %2d", i)
		err = rmq.PublishWithDelay("user.event.publish", []byte(msg), int64(1000*i))
		if err != nil {
			log.Fatalf("run: failed to publish into rabbitmq: %v", err)
		}
		time.Sleep(100 * time.Millisecond)
	}

	select {}
}
