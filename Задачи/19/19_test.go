package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	validTestData := []struct {
		s        string
		expected string
	}{
		{
			s:        "главрыба",
			expected: "абырвалг",
		},
		{
			s:        "12345фq ",
			expected: " qф54321",
		},
		{
			s:        "м",
			expected: "м",
		},
		{
			s:        "",
			expected: "",
		},
	}

	for _, data := range validTestData {
		res := ReverseString(data.s)
		assert.Equal(t, data.expected, res)
	}
}
