package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSubset(t *testing.T) {
	validTestData := []struct {
		arr      []float64
		expected map[int][]float64
	}{
		{
			arr: []float64{-10, -7, 0, 6.7, 10, 20, 25, 34},
			expected: map[int][]float64{
				-10: {-10},
				-1:  {-7},
				0:   {0, 6.7},
				10:  {10},
				20:  {20, 25},
				30:  {34},
			},
		},
		{
			arr: []float64{-25.4, -27, 13, 19, 15.5, 24.5, -21, 32.5},
			expected: map[int][]float64{
				-20: {-25.4, -27, -21},
				10:  {13, 19, 15.5},
				20:  {24.5},
				30:  {32.5},
			},
		},
	}

	for _, data := range validTestData {
		res := createSubset(data.arr)
		assert.Equal(t, data.expected, res)
	}
}
