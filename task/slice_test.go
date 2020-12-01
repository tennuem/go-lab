package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Add добавляет число в переданный слайс.
// Создает копию структуры слайса и работает с ней, но c тем же блоком памяти.
func Add(arr []int, v int) {
	// arr{ptr: 0xc00000e440, len: 2, cap: 4}
	arr = append(arr, v)
	// arr{ptr: 0xc00000e440, len: 3, cap: 4}
}

func TestAdd(t *testing.T) {
	arr := make([]int, 2, 4)
	arr[0] = 1
	arr[1] = 3

	// arr{ptr: 0xc00000e400, len: 2, cap: 4}
	Add(arr, 5)
	// arr{ptr: 0xc00000e400, len: 2, cap: 4}

	// все элементы
	assert.Equal(t, []int{1, 3}, arr)

	// первые 3 элемента
	assert.Equal(t, []int{1, 3, 5}, arr[:3])

	// перезапишет 3 число, тк длина слайса равна 2
	arr = append(arr, 8)
	// arr{ptr: 0xc00000e400, len: 3, cap: 4}
	assert.Equal(t, []int{1, 3, 8}, arr)
}
