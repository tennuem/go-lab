package task

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func palindrome(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-i-1] {
			return false
		}
	}
	return true
}

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		in  string
		out bool
	}{
		{"rotor", true},
		{"asd", false},
		{"ротор", true},
	}
	for i, c := range testCases {
		assert.Equal(t, c.out, palindrome(c.in), fmt.Sprintf("case-%d", i))
	}
}
