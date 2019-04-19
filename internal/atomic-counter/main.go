package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	var ac AtomicCounter
	for i := 0; i < 100; i++ {
		go func(no int) {
			for i := 0; i < 10000; i++ {
				ac.Add(1)
			}
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(ac.Value())
}

type AtomicCounter struct {
	Val int64
}

func (ac *AtomicCounter) Add(n int64) {
	atomic.AddInt64(&ac.Val, n)
	runtime.Gosched()
}

func (ac *AtomicCounter) Value() (n int64) {
	return atomic.LoadInt64(&ac.Val)
}
