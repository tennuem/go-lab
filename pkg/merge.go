package pkg

import (
	"sync"
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
