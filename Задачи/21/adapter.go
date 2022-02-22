package main

import "math/rand"

// Adapter - адаптер, который адаптирует структуру Sensor под интерфейс SomeDeviceInterface без изменения логики
// самой структуры Sensor. Данные, которые получаются с помощью метода Sensor.GetFahrenheitTemperature, адаптер
// должен преобразовать из градусов по Фаренгейту в градусы Цельсия, для того чтобы с этими покзаниями
// могла работать функция DoSomeWithCelsiusTemperature().
// Таким образом, Sensor - структура, которую нужно адаптировать.
// А SomeDeviceInterface - интерфейс, под который адаптер должен адаптировать структуру Sensor.
type Adapter struct {
	sensor Sensor
}

// GetCelsiusTemperature - переводит показания из грудусов по Фаренгейту в градусы Цельсия.
func (a Adapter) GetCelsiusTemperature() float64 {
	return a.sensor.GetFahrenheitTemperature()*9/5 + 32
}

// DoSomeWithCelsiusTemperature - работает с показаниями Sensor и некоторым образом преобразовывает их,
// возвращая затем значение.
func (a Adapter) DoSomeWithCelsiusTemperature() float64 {
	return a.GetCelsiusTemperature() * rand.Float64() / rand.Float64()
}
