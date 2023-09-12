package play031_timer

import (
	"fmt"
	"time"
)

func Play() {
	exitChan := make(chan struct{})

	go func() {
		for {
			select {
			case <-exitChan:
				fmt.Println("协程收到退出信号，退出")
				return
			default:
				// 协程的主要工作
				fmt.Println("协程正在执行...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 运行一段时间后发送退出信号
	time.Sleep(5 * time.Second)
	close(exitChan)
	time.Sleep(1 * time.Second)
	fmt.Println("主程序退出")
}
