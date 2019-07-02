package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorker(t *testing.T) {
	var (
		workers = 4
		jobs    = 10
		jobCh   = make(chan int, 100)
		resCh   = make(chan int, 100)
	)

	for i := 0; i < workers; i++ {
		go Worker(i, jobCh, resCh)
	}

	for j := 0; j < jobs; j++ {
		jobCh <- j
	}
	close(jobCh)

	for i := 0; i < jobs; i++ {
		res := <-resCh

		assert.Equal(t, res/2, i)
	}
}
