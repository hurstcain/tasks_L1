package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuicksort(t *testing.T) {
	validTestData := []struct {
		arr      []int
		expected []int
	}{
		{
			arr:      []int{1, 0, 200, -19, 3, 300},
			expected: []int{-19, 0, 1, 3, 200, 300},
		},
		{
			arr:      []int{-20, -30, 10, 1, 3},
			expected: []int{-30, -20, 1, 3, 10},
		},
		{
			arr:      []int{20},
			expected: []int{20},
		},
		{
			arr:      []int{},
			expected: []int{},
		},
	}

	for _, data := range validTestData {
		res := data.arr
		Quicksort(res, 0, len(res)-1)
		assert.Equal(t, data.expected, res)
	}
}
