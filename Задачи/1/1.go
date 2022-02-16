package main

import "fmt"

type Human struct {
	name string
	age  uint
	sex  string
}

func (h Human) Name() string {
	return h.name
}

func (h Human) Age() uint {
	return h.age
}

func (h Human) Sex() string {
	return h.sex
}

func (h Human) FullInfo() string {
	return fmt.Sprintf("Name: %s, age: %d, sex: %s\n", h.name, h.age, h.sex)
}

type Action struct {
	// Внедряем структуру Human в Action.
	// После этого все методы Human будут доступны через Action.
	Human
}
