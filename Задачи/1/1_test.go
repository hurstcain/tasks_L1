package main

import (
	"testing"
)

var act = Action{
	Human{
		name: "Alina",
		age:  24,
		sex:  "female",
	},
}

func TestName(t *testing.T) {
	result := act.Name()
	expected := "Alina"

	t.Logf("TestName. Result: %v", result)

	if result != expected {
		t.Errorf("Incorrect result. Expected %v, but got %v.", expected, result)
	} else {
		t.Log("Test passed")
	}
}

func TestAge(t *testing.T) {
	result := act.Age()
	var expected uint = 24

	t.Logf("TestAge. Result: %v", result)

	if result != expected {
		t.Errorf("Incorrect result. Expected %v, but got %v.", expected, result)
	} else {
		t.Log("Test passed")
	}
}

func TestSex(t *testing.T) {
	result := act.Sex()
	expected := "female"

	t.Logf("TestSex. Result: %v", result)

	if result != expected {
		t.Errorf("Incorrect result. Expected %v, but got %v.", expected, result)
	} else {
		t.Log("Test passed")
	}
}

func TestFullInfo(t *testing.T) {
	result := act.FullInfo()
	expected := "Name: Alina, age: 24, sex: female\n"

	t.Logf("TestFullInfo. Result: %v", result)

	if result != expected {
		t.Errorf("Incorrect result. Expected %v, but got %v.", expected, result)
	} else {
		t.Log("Test passed")
	}
}
