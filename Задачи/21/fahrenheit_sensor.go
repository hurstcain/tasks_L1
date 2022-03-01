package main

import (
	"math/rand"
)

type Fahrenheit float64

type FahrenheitSensor interface {
	GetFahrenheitTemperature() Fahrenheit
}

// FSensor - структура, которая реализует интерфейс FahrenheitSensor.
// Представляет собой имитацию датчика, который измеряет температуру в градусах по Фаренгейту.
type FSensor struct{}

// GetFahrenheitTemperature - возвращает значение температуры в градусах по Фаренгейту.
func (f *FSensor) GetFahrenheitTemperature() Fahrenheit {
	return Fahrenheit(rand.Float64()*float64(rand.Intn(11)) + float64(rand.Intn(50)))
}
