package main

// Adapter - адаптер, который адаптирует структуру FSensor под интерфейс CelsiusSensor без изменения логики
// самой структуры FSensor. Данные, которые получаются с помощью метода FSensor.GetFahrenheitTemperature, адаптер
// должен преобразовать из градусов по Фаренгейту в градусы Цельсия для того, чтобы с этими показаниями
// могла работать функция DoSomeWithCelsiusTemperature().
// Таким образом, FSensor - структура, которую нужно адаптировать.
// А CelsiusSensor - интерфейс, под который адаптер должен адаптировать структуру Sensor.
type Adapter struct {
	sensor FSensor
}

// GetCelsiusTemperature - переводит показания из градусов по Фаренгейту в градусы Цельсия.
func (a Adapter) GetCelsiusTemperature() Celsius {
	return (Celsius(a.sensor.GetFahrenheitTemperature()) - 32) * 5 / 9
}
