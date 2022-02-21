package main

import (
	"math"
)

// Point - структура, содержащая две переменные x и y, которые представляют собой координаты точки.
type Point struct {
	x, y float64
}

// NewPoint - конструктор структуры Point
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// DistanceBetweenTwoPoints - находит расстояние между двумя точками a и b
func DistanceBetweenTwoPoints(a, b Point) float64 {
	return math.Sqrt(math.Pow(b.x-a.x, 2) + math.Pow(b.y-a.y, 2))
}
