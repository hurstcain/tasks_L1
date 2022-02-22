package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckBigintValue(t *testing.T) {
	validTestData := []struct {
		a bigint
	}{
		{
			a: 1<<21 + 676,
		},
		{
			a: 1 << 34,
		},
	}

	errorTestData := []struct {
		a bigint
	}{
		{
			a: 1 << 20,
		},
		{
			a: -10,
		},
	}

	for _, data := range validTestData {
		err := checkBigintValue(data.a)
		assert.NoError(t, err)
	}

	for _, data := range errorTestData {
		err := checkBigintValue(data.a)
		assert.Error(t, err)
	}
}

func TestMult(t *testing.T) {
	validTestData := []struct {
		a        bigint
		b        bigint
		expected int64
	}{
		{
			a:        1 << 21,
			b:        1 << 21,
			expected: 1 << 42,
		},
		{
			a:        1<<30 - 6762,
			b:        1<<23 + 9229,
			expected: 9017052031860894,
		},
	}

	errorTestData := []struct {
		a bigint
		b bigint
	}{
		{
			a: 1 << 60,
			b: 1 << 60,
		},
	}

	for _, data := range validTestData {
		res, err := mult(data.a, data.b)
		assert.Equal(t, res, data.expected)
		assert.NoError(t, err)
	}

	for _, data := range errorTestData {
		_, err := mult(data.a, data.b)
		assert.Error(t, err)
	}
}

func TestDiv(t *testing.T) {
	validTestData := []struct {
		a        bigint
		b        bigint
		expected int64
	}{
		{
			a:        1 << 21,
			b:        1 << 21,
			expected: 1,
		},
		{
			a:        1<<30 - 6762,
			b:        1<<23 + 9229,
			expected: 127,
		},
		{
			a:        1 << 23,
			b:        1 << 40,
			expected: 0,
		},
	}

	for _, data := range validTestData {
		res := div(data.a, data.b)
		assert.Equal(t, res, data.expected)
	}
}

func TestSum(t *testing.T) {
	validTestData := []struct {
		a        bigint
		b        bigint
		expected int64
	}{
		{
			a:        1 << 21,
			b:        1 << 21,
			expected: 4194304,
		},
		{
			a:        1<<30 - 6762,
			b:        1<<23 + 9229,
			expected: 1082132899,
		},
	}

	errorTestData := []struct {
		a bigint
		b bigint
	}{
		{
			a: 1 << 62,
			b: 1 << 62,
		},
	}

	for _, data := range validTestData {
		res, err := sum(data.a, data.b)
		assert.Equal(t, res, data.expected)
		assert.NoError(t, err)
	}

	for _, data := range errorTestData {
		_, err := sum(data.a, data.b)
		assert.Error(t, err)
	}
}

func TestSub(t *testing.T) {
	validTestData := []struct {
		a        bigint
		b        bigint
		expected int64
	}{
		{
			a:        1 << 21,
			b:        1 << 21,
			expected: 0,
		},
		{
			a:        1<<30 - 6762,
			b:        1<<23 + 9229,
			expected: 1065337225,
		},
		{
			a:        1<<23 - 1,
			b:        1 << 25,
			expected: -25165825,
		},
	}

	for _, data := range validTestData {
		res := sub(data.a, data.b)
		assert.Equal(t, res, data.expected)
	}
}
