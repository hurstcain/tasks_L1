package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersection(t *testing.T) {
	validTestData := []struct {
		a        []interface{}
		b        []interface{}
		expected []interface{}
	}{
		{
			a:        []interface{}{1, 3, 0, "set", 20, "0", 21},
			b:        []interface{}{1, "set", "0", 21, 0, 73, 902, "1", "3"},
			expected: []interface{}{1, "set", "0", 21, 0},
		},
		{
			a:        []interface{}{1, 2, 3},
			b:        []interface{}{},
			expected: []interface{}{},
		},
		{
			a:        []interface{}{},
			b:        []interface{}{2, 40, "a"},
			expected: []interface{}{},
		},
		{
			a:        []interface{}{1, "b", 30},
			b:        []interface{}{2, 40, "a"},
			expected: []interface{}{},
		},
	}

	for _, data := range validTestData {
		res := intersection(data.a, data.b)
		assert.Equal(t, data.expected, res)
	}
}
