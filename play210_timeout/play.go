package play210_timeout

import (
	"context"
	"log"
	"time"
)

func process(ctx context.Context, data string) {
	// 模拟一些耗时操作
	select {
	case <-time.After(5 * time.Second):
		log.Println(data, "processed successfully")
	case <-ctx.Done():
		log.Println(data, "processing canceled:", ctx.Err())
	}

	log.Println("done ")
}

func Play() {
	// 创建一个父上下文
	parentCtx := context.Background()

	// 创建一个带有取消功能的子上下文，设置超时时间为3秒
	childCtx, cancel := context.WithTimeout(parentCtx, 3*time.Second)
	defer cancel()

	// 启动goroutine来处理数据
	log.Println("starting")
	go process(childCtx, "Data 1")
	go process(childCtx, "Data 2")

	// 等待一段时间，以便观察取消效果
	select {}
}
