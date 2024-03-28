package main

import (
	"20240328094701/redislock"
	"log"
)

func main() {
	e := redislock.GetInstance().Init("192.168.31.7:6379", "lock-name")
	if e != nil {
		log.Printf("GetMqInstance().Init e = %#v", e)
		return
	}

	var cnt = 0
	for i := 0; i < 200; i++ {
		go func(tmp int) {
			if err := redislock.GetInstance().Lock(); err != nil {
				log.Printf("Lock err = %s \n", err)
			}
			//
			//
			cnt++
			log.Printf("biz logic id: %3d , cnt: %3d \n", tmp, cnt)
			//
			//
			//cnt--
			if ok, err := redislock.GetInstance().Unlock(); !ok || err != nil {
				log.Printf("Unlock err = %s \n", err)
			}
		}(i)
	}

	select {}

	return
}
