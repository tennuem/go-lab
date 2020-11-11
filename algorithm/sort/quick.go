package sort

// QuickSort https://en.wikipedia.org/wiki/Quicksort
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	var (
		median = arr[0]
		low    = make([]int, 0, len(arr))
		medium = make([]int, 0, len(arr))
		high   = make([]int, 0, len(arr))
	)

	for _, item := range arr {
		switch {
		case item < median:
			low = append(low, item)
		case item == median:
			medium = append(medium, item)
		case item > median:
			high = append(high, item)
		}
	}

	low = QuickSort(low)
	high = QuickSort(high)

	low = append(low, medium...)
	low = append(low, high...)

	return low
}
