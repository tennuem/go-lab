package task

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Merge converts a list of channels to a single channel by starting a goroutine
// for each inbound channel that copies the values to the sole outbound channel.
// Once all the output goroutines have been started,
// merge starts one more goroutine to close the outbound channel after all sends on that channel are done.
func Merge(quit <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				select {
				case out <- n:
				case <-quit:
					return
				}
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

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
