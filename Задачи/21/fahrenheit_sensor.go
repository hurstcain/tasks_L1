package main

import (
	"math/rand"
)

type FahrenheitSensor interface {
	GetFahrenheitTemperature() float64
}

// Sensor - структура, которая реализует интерфейс FahrenheitSensor.
// Представляет собой имитацию датчика, который измеряет температуру в градусах по Фаренгейту.
type Sensor struct{}

// GetFahrenheitTemperature - возвращает значение температуры в градусах по Фаренгейту.
func (s *Sensor) GetFahrenheitTemperature() float64 {
	return rand.Float64()
}
