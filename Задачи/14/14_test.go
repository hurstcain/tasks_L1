package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefineType1(t *testing.T) {
	validTestData := []struct {
		a        interface{}
		expected string
	}{
		{
			a:        10,
			expected: "int",
		},
		{
			a:        "1234",
			expected: "string",
		},
		{
			a:        false,
			expected: "bool",
		},
	}

	for _, data := range validTestData {
		res := defineType1(data.a)
		assert.Equal(t, data.expected, res)
	}
}

func TestDefineType2(t *testing.T) {
	validTestData := []struct {
		a        interface{}
		expected string
	}{
		{
			a:        10,
			expected: "int",
		},
		{
			a:        "1234",
			expected: "string",
		},
		{
			a:        false,
			expected: "bool",
		},
	}

	for _, data := range validTestData {
		res := defineType2(data.a)
		assert.Equal(t, data.expected, res)
	}
}

func TestDefineType3(t *testing.T) {
	validTestData := []struct {
		a        interface{}
		expected string
	}{
		{
			a:        10,
			expected: "int",
		},
		{
			a:        "1234",
			expected: "string",
		},
		{
			a:        false,
			expected: "bool",
		},
	}

	invalidTestData := []struct {
		a        interface{}
		expected string
	}{
		{
			a:        11.006,
			expected: "",
		},
	}

	for _, data := range validTestData {
		res, err := defineType3(data.a)
		assert.Equal(t, data.expected, res)
		assert.NoError(t, err)
	}

	for _, data := range invalidTestData {
		res, err := defineType3(data.a)
		assert.Equal(t, data.expected, res)
		assert.Error(t, err)
	}
}
