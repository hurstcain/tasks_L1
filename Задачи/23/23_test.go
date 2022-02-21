package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteElement(t *testing.T) {
	validTestData := []struct {
		arr      []int
		i        int
		expected []int
	}{
		{
			arr:      []int{1, 2, 3, 4, 5},
			i:        0,
			expected: []int{2, 3, 4, 5},
		},
		{
			arr:      []int{1, 2, 3, 4, 5},
			i:        4,
			expected: []int{1, 2, 3, 4},
		},
		{
			arr:      []int{1, 2, 3, 4, 5},
			i:        3,
			expected: []int{1, 2, 3, 5},
		},
	}

	errorTestData := []struct {
		arr []int
		i   int
	}{
		{
			arr: []int{1, 2, 3, 4, 5},
			i:   100,
		},
		{
			arr: []int{1, 2, 3, 4, 5},
			i:   -1,
		},
	}

	for _, data := range validTestData {
		err := DeleteElement(&data.arr, data.i)
		assert.Equal(t, data.expected, data.arr)
		assert.NoError(t, err)
	}

	for _, data := range errorTestData {
		err := DeleteElement(&data.arr, data.i)
		assert.Error(t, err)
	}
}
