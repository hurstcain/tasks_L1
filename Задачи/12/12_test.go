package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSet(t *testing.T) {
	validTestData := []struct {
		a        []string
		expected []string
	}{
		{
			a:        []string{"cat", "cat", "dog", "cat", "tree"},
			expected: []string{"cat", "dog", "tree"},
		},
		{
			a:        []string{},
			expected: []string{},
		},
		{
			a:        []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
	}

	for _, data := range validTestData {
		res := createSet(data.a)
		assert.Equal(t, len(data.expected), len(res))
	}
}
