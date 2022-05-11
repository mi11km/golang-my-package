package utils

import (
	"fmt"
	"time"
)

func generator(done <-chan interface{}) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 0; i < 10; i++ {
			select {
			case <-done:
				return
			default:
				time.Sleep(time.Second * 1)
				intStream <- i
			}
		}
	}()
	return intStream
}

func consumer(done chan interface{}, stream <-chan int) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println(<-stream)
		}
	}

}
