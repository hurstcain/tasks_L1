package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	validTestData := []struct {
		arr      []int
		el       int
		expected int
	}{
		{
			arr:      []int{-10, 0, 12, 45, 98, 10},
			el:       0,
			expected: 1,
		},
		{
			arr:      []int{1, 2, 5, 7, 9, 10, 23},
			el:       10,
			expected: 5,
		},
		{
			arr:      []int{8},
			el:       8,
			expected: 0,
		},
		{
			arr:      []int{},
			el:       10,
			expected: -1,
		},
		{
			arr:      []int{20, 34, 90, 103, 500, 3000},
			el:       0,
			expected: -1,
		},
	}

	for _, data := range validTestData {
		res := BinarySearch(data.arr, data.el)
		assert.Equal(t, data.expected, res)
	}
}
