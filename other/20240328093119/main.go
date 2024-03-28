package main

import (
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"log"
)

/*
go get github.com/go-redis/redis/v8
go get "github.com/go-redsync/redsync/v4"
go get "github.com/go-redsync/redsync/v4/redis/goredis/v8"
*/

func main() {
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "192.168.31.7:6379",
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)
	rs := redsync.New(pool)
	mutex := rs.NewMutex("my-global-mutex")
	log.Printf("mutex: %#v", mutex)

	var cnt = 0
	for i := 0; i < 20; i++ {
		//time.Sleep(time.Millisecond)
		go func(tmp int) {
			if err := mutex.Lock(); err != nil {
				log.Printf("Lock err = %s \n", err)
			}
			cnt++
			log.Printf("biz logic id: %3d , cnt: %3d \n", tmp, cnt)

			//cnt--
			if ok, err := mutex.Unlock(); !ok || err != nil {
				log.Printf("Unlock err = %s \n", err)
			}
		}(i)
	}

	select {}
}
