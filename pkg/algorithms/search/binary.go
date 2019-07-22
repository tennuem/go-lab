package search

// Binary https://en.wikipedia.org/wiki/Binary_search_algorithm
func Binary(arr []int, n int) bool {
	var (
		left  = 0
		right = len(arr) - 1
	)
	for left <= right {
		mid := (left + right) / 2
		switch {
		case arr[mid] < n:
			left = mid + 1
		case arr[mid] > n:
			right = mid - 1
		default:
			return true
		}
	}
	return false
}
