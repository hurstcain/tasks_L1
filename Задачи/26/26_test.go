package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAreSymbolsUnique(t *testing.T) {
	validTestData := []struct {
		s        string
		expected bool
	}{
		{
			s:        "abcd",
			expected: true,
		},
		{
			s:        "abCdefAaf",
			expected: false,
		},
		{
			s:        "aabcd",
			expected: false,
		},
		{
			s:        "abcdA",
			expected: false,
		},
		{
			s:        "абвгд",
			expected: true,
		},
		{
			s:        "аб124БА",
			expected: false,
		},
		{
			s:        "12345",
			expected: true,
		},
		{
			s:        "ф133",
			expected: false,
		},
	}

	for _, data := range validTestData {
		res := areSymbolsUnique(data.s)
		assert.Equal(t, data.expected, res)
	}
}
