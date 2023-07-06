package play012_gopark

import (
	"runtime"
	"sync"
	"sync/atomic"
)

func Play() {
	var wg sync.WaitGroup
	var counter uint64

	wg.Add(2)

	go func() {
		defer wg.Done()

		// 让goroutine进入休眠状态
		runtime.Gosched()
		atomic.AddUint64(&counter, 1)
	}()

	//go func() {
	//	defer wg.Done()
	//
	//	// 让goroutine进入休眠状态，并在条件满足时被唤醒
	//	runtime.gopark(func(_ unsafe.Pointer, _ unsafe.Pointer) bool {
	//		// 假设条件是counter的值达到10
	//		return atomic.LoadUint64(&counter) >= 10
	//	}, nil, "customUnblocker", 0)
	//}()

	wg.Wait()
}
