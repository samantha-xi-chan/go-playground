package play901_pprof

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

func Play() {
	// 故意生成 100个 协程
	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(time.Second * 3600)
		}()
	}

	for i := 0; i < 2; i++ {
		go func() {
			for {
				i = i + 1
				//log.Println(i)
			}
		}()
	}

	log.Println("time.Tick ing222")
	// 开启pprof，监听请求 http://localhost:6060/debug/pprof/
	ip := "0.0.0.0:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}

	log.Println("time.Tick ing")

	tick := time.Tick(time.Second / 100)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}
}

func PlayPprofLocalfile() {
	f, _ := os.OpenFile("cpu.profile", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	n := 10
	for i := 1; i <= 5; i++ {
		fmt.Printf("fib(%d)=%d\n", n, fib(n))
		n += 3 * i
	}
}

func fib(n int) int {
	if n <= 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func InitPProf() {
	// 开启pprof，监听请求 http://localhost:6060/debug/pprof/
	ip := "0.0.0.0:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}
