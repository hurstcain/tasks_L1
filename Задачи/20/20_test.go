package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWords(t *testing.T) {
	validTestData := []struct {
		s        string
		expected string
	}{
		{
			s:        "snow dog sun",
			expected: "sun dog snow",
		},
		{
			s:        "1 2 3 4",
			expected: "4 3 2 1",
		},
		{
			s:        "snow",
			expected: "snow",
		},
		{
			s:        "",
			expected: "",
		},
	}

	for _, data := range validTestData {
		res := ReverseWords(data.s)
		assert.Equal(t, data.expected, res)
	}
}
