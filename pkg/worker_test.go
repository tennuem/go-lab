package pkg

import (
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorker(t *testing.T) {
	var (
		wg      sync.WaitGroup
		workers = 4
		jobs    = 10
		jobCh   = make(chan int, 2)
		resCh   = make(chan int, 2)
	)

	// send to worker
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			close(jobCh)
		}()
		for j := 0; j < jobs; j++ {
			jobCh <- j
		}
	}()

	// get from worker
	var result []int
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			close(resCh)
		}()
		for j := 0; j < jobs; j++ {
			result = append(result, <-resCh)
		}
	}()

	for i := 0; i < workers; i++ {
		go Worker(i, jobCh, resCh)
	}

	wg.Wait()

	sort.Ints(result)
	for i := 0; i < jobs; i++ {
		assert.Equal(t, result[i]/2, i)
	}
}
