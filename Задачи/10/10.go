package main

import "fmt"

// Функция, которая создает подмножества.
// Подмножества хранятся в мапе по ключам.
func createSubset(arr []float64) map[int][]float64 {
	m := make(map[int][]float64)

	for _, val := range arr {
		if val < 0 && int(val/10) == 0 {
			key := -1
			m[key] = append(m[key], val)
			continue
		}
		key := int(val/10) * 10
		m[key] = append(m[key], val)
	}

	return m
}

func main() {
	arr := []float64{
		-25.4, -27, 13, 19, 15.5, 24.5, -21, 32.5,
	}

	m := createSubset(arr)

	fmt.Println(m)
}
