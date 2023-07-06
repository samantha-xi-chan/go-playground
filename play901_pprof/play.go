package play901_pprof

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func Play() {

	// 故意生成 100个 协程
	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(time.Second * 3600)
		}()
	}

	// 开启pprof，监听请求 http://localhost:6060/debug/pprof/
	ip := "0.0.0.0:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}

	tick := time.Tick(time.Second / 100)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}
}
