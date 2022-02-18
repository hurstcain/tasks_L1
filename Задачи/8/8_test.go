package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseBit(t *testing.T) {
	validTestData := []struct {
		a        int64
		bit      int
		expected int64
	}{
		{
			0, 2, 2,
		},
		{
			256, 9, 0,
		},
		{
			256, 1, 257,
		},
		{
			-10, 3, -14,
		},
		{
			-1024, 11, 0,
		},
	}

	errorTestData := []struct {
		a   int64
		bit int
	}{
		{
			3, -100,
		},
		{
			3, 10000,
		},
	}

	for _, data := range validTestData {
		res, _ := reverseBit(data.a, data.bit)
		assert.Equal(t, data.expected, res)
	}

	for _, data := range errorTestData {
		_, err := reverseBit(data.a, data.bit)
		assert.Error(t, err)
	}
}
