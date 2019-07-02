package sort

import (
	"time"
)

func NewMerge() Sort {
	return &mergeSort{}
}

type mergeSort struct {
	Items      []int
	Operations int
	Durations  time.Duration
}

func (s *mergeSort) Result(array []int) []int {

	// Time of algorithm execution
	start := time.Now()
	defer func(start time.Time) {
		s.Durations = time.Since(start)
	}(start)

	result := sort(array)

	s.Operations = count

	return result
}

func (s *mergeSort) Operation() int {
	return s.Operations
}

func (s *mergeSort) Duration() time.Duration {
	return s.Durations
}

var count int

func merge(left, right []int) []int {
	ret := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(ret, right...)
		}
		if len(right) == 0 {
			return append(ret, left...)
		}
		if left[0] <= right[0] {
			ret = append(ret, left[0])
			left = left[1:]
		} else {
			ret = append(ret, right[0])
			right = right[1:]
		}
	}
	return ret
}

func sort(s []int) []int {

	count++

	if len(s) <= 1 {
		return s
	}

	middle := len(s) / 2
	left := sort(s[:middle])
	right := sort(s[middle:])
	return merge(left, right)
}
