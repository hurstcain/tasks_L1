package main

import "math/rand"

type Celsius float64

type CelsiusSensor interface {
	GetCelsiusTemperature() Celsius
}

// CSensor - структура, реализующая интерфейс CelsiusSensor.
// Представляет собой имитацию устройства, которое измеряет температуру в градусах Цельсия.
type CSensor struct{}

// GetCelsiusTemperature - возвращает значение температуры в градусах Цельсия.
func (c *CSensor) GetCelsiusTemperature() Celsius {
	return Celsius(rand.Float64())
}

// DoSomeWithCelsiusTemperature - производит некоторые операции со значением градусов Цельсия
// и возвращает получившийся результат.
func DoSomeWithCelsiusTemperature(d Celsius) float64 {
	return float64(d) * rand.Float64() / rand.Float64()
}
