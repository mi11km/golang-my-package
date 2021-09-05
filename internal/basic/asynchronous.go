package basic

import (
	"fmt"
	"runtime"
	"sync"
)

func Test() {
	var wg sync.WaitGroup
	for _, word := range []string{"hello", "world", "hooooo"} {
		wg.Add(1)
		// 参照だから hooooo しか表示されない
		go func() {
			defer wg.Done()
			fmt.Println(word)
		}()

		// コピーを渡してるので、すべての文字がちゃんと表示される
		//go func(w string) {
		//	defer wg.Done()
		//	fmt.Println(word)
		//}(word)
	}
	wg.Wait()
}

func MeasureGroutineSize() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Println(float64(after-before)/numGoroutines/1000, " kb")
}
