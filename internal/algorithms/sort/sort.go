package sort

import "time"

type Sort interface {
	Result(array []int) []int
	Operation() int
	Duration() time.Duration
}
