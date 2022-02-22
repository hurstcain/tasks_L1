package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	adapter := Adapter{
		sensor: Sensor{},
	}

	fmt.Printf("Показания в градусах по Фаренгейту: %v\n", adapter.sensor.GetFahrenheitTemperature())
	fmt.Printf("Показания в градусах Цельсия: %v\n", adapter.GetCelsiusTemperature())
	fmt.Printf("Преобразование показаний: %v\n", adapter.DoSomeWithCelsiusTemperature())
}
