package main

import "math/rand"

type SomeDeviceInterface interface {
	GetCelsiusTemperature() float64
	DoSomeWithCelsiusTemperature() float64
}

// CelsiusTemperature - структура, реализующая интерфейс SomeDeviceInterface.
// Представляет собой имитацию устройства, которое получает значение температуры в градусах, а затем
// выполняет со значениями некоторые преобразования.
type CelsiusTemperature struct{}

// GetCelsiusTemperature - возвращает значение температуры в градусах.
func (ct *CelsiusTemperature) GetCelsiusTemperature() float64 {
	return rand.Float64()
}

// DoSomeWithCelsiusTemperature - производит некоторые операции со значением, возвращаемым методом
// GetCelsiusTemperature, и возвращет получившийся результат.
func (ct CelsiusTemperature) DoSomeWithCelsiusTemperature() float64 {
	return ct.GetCelsiusTemperature() * rand.Float64() / rand.Float64()
}
