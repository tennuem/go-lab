package pkg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAtomicCounter(t *testing.T) {
	var ac AtomicCounter
	for i := 0; i < 100; i++ {
		go func(no int) {
			for i := 0; i < 10000; i++ {
				ac.Add(1)
			}
		}(i)
	}
	time.Sleep(time.Second)

	assert.Equal(t, int(ac.Value()), 1000000)
}
