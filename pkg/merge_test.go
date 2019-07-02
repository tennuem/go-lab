package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	quit := make(chan struct{})
	defer close(quit)

	ch1 := write(quit)
	ch2 := write(quit)

	var count int
	for _ = range Merge(quit, ch1, ch2) {
		count++
	}

	assert.Equal(t, count, 20)
}

// write pushes 10 ints to chan
func write(quit <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			select {
			case out <- i:
			case <-quit:
				return
			}
		}
	}()

	return out
}
