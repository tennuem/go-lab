package pkg

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomicCounter(t *testing.T) {
	var (
		wg sync.WaitGroup
		ac AtomicCounter
	)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(no int) {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				ac.Add(1)
			}
		}(i)
	}
	wg.Wait()

	assert.Equal(t, int(ac.Value()), 1000000)
}
