package main

import "fmt"

// Меняет местами значения двух переменных a и b
func swap(a, b *int) {
	*a, *b = *b, *a
}

// Меняет местами значения двух переменных a и b с помощью операций вычитания и сложения
func swapArithmetic(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func main() {
	a, b := 10, 0
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
	swapArithmetic(&a, &b)
	fmt.Println(a, b)
}
