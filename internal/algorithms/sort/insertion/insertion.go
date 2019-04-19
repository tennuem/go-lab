package insertion

import (
	"time"

	"github.com/tennuem/go-lab/internal/algorithms/sort"
)

type Insertion struct {
	Operations int
	Durations  time.Duration
}

func (ins *Insertion) Result(array []int) []int {

	// Time of algorithm execution
	start := time.Now()
	defer func(start time.Time) {
		ins.Durations = time.Since(start)
	}(start)

	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {

			// Counting number operations
			ins.Operations++

			array[j-1], array[j] = array[j], array[j-1]
		}
	}
	return array
}

func (ins *Insertion) Operation() int {
	return ins.Operations
}

func (ins *Insertion) Duration() time.Duration {
	return ins.Durations
}

func NewInsertion() sort.Sort {
	return &Insertion{}
}
