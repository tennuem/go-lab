package task

import (
	"testing"
)

// matrix вычисляет сумму диагоналей матрицы
func matrix(m [][]int) int {
	var sum int
	for i := 0; i < len(m); i++ {
		sum += m[i][i]
		sum += m[i][len(m)-1-i]
	}
	if len(m)%2 != 0 {
		sum -= m[len(m)/2][len(m)/2]
	}
	return sum
}

func TestMatrix(t *testing.T) {
	case1 := [][]int{
		{1, 2},
		{3, 4},
	}
	case2 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	if 10 != matrix(case1) {
		t.Error("case-1 not equal 10")
	}
	if 25 != matrix(case2) {
		t.Error("case-2 not equal 25")
	}
}
