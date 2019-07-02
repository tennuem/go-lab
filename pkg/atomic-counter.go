package pkg

import (
	"runtime"
	"sync/atomic"
)

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
