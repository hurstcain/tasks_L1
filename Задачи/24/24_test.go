package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestDistanceBetweenTwoPoints(t *testing.T) {
	validTestData := []struct {
		a        Point
		b        Point
		expected float64
	}{
		{
			a:        NewPoint(10, 20),
			b:        NewPoint(30, 5),
			expected: 25,
		},
		{
			a:        NewPoint(5.54, 1.8),
			b:        NewPoint(-7.1, 9),
			expected: 14.5,
		},
		{
			a:        NewPoint(-20, -3),
			b:        NewPoint(-6.98, 0),
			expected: 13.4,
		},
		{
			a:        NewPoint(0, 0),
			b:        NewPoint(0, 0),
			expected: 0,
		},
	}

	for _, data := range validTestData {
		res := DistanceBetweenTwoPoints(data.a, data.b)
		res = math.Round(res*10) / 10
		assert.Equal(t, data.expected, res)
	}
}
