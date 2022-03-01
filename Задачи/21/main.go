package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	adapter := Adapter{
		sensor: FSensor{},
	}

	c := adapter.GetCelsiusTemperature()

	fmt.Printf("Показания в градусах Цельсия: %v\n", c)
	fmt.Printf("Преобразование показаний: %v\n", DoSomeWithCelsiusTemperature(c))
	// output:
	// Показания в градусах Цельсия: 8.10163994066513
	// Преобразование показаний: 23.30759359726299
}
