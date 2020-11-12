package task

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func fanIn(channels ...chan int) chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	go func() {
		wg.Wait()
		close(out)
	}()
	for _, ch := range channels {
		wg.Add(1)
		go func(ch chan int) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(ch)
	}
	return out
}

func TestFanIn(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	out := fanIn(ch1, ch2)
	ch1 <- 1
	ch2 <- 2
	close(ch1)
	close(ch2)
	var res []int
	for v := range out {
		res = append(res, v)
	}
	assert.Equal(t, []int{1, 2}, res)
}
