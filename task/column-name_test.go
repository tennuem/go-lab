package task

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var alphabet = []rune{'A', 'B', 'C'}

// return column name (like excel):
// 0 - A
// 1 - B
// 2 - C
// 3 - AA
// 4 - AB
// ...
func columnName(index int) string {
	a := index / 3
	b := index % 3
	if index > 2 {
		index = 0
	}
	if a > 0 {
		return string(alphabet[a-1]) + string(alphabet[b])
	}
	return string(alphabet[index])
}

func TestColumnName(t *testing.T) {
	testCases := []string{"A", "B", "C", "AA", "AB", "AC", "BA", "BB", "BC", "CA", "CB", "CC"}
	for i, c := range testCases {
		assert.Equal(t, c, columnName(i), fmt.Sprintf("case-%d", i))
	}
}
