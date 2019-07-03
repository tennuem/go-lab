package sort

//InsertionSort1 https://en.wikipedia.org/wiki/Insertion_sort
func InsertionSort1(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr
}

//InsertionSort2 https://en.wikipedia.org/wiki/Insertion_sort
func InsertionSort2(arr []int) []int {
	var i, j int
	for i = 1; i < len(arr); i++ {
		for j = 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
