package main

import (
	"errors"
	"fmt"
	"log"
)

// Тип больших чисел, которые больше 2^20
type bigint int64

// Проверка, соответствует ли значение переменной a типу bigint
func checkBigintValue(a bigint) error {
	if a > 1<<20 {
		return nil
	}

	return errors.New("bigint is too small")
}

// Функция умножения больших чисел. Если результат умножения вышел за пределы максимального значения типа int64,
// то возвращается ошибка
func mult(a, b bigint) (int64, error) {
	res := int64(a * b)
	if a > 0 && b > 0 && res <= 0 {
		return 0, errors.New("result of multiplication is out of the range of type int64")
	}

	return res, nil
}

// Функция деления больших чисел
func div(a, b bigint) int64 {
	return int64(a / b)
}

// Функция суммирования больших чисел. Если сумма вышла за пределы максимального значения типа int64,
// то возвращается ошибка
func sum(a, b bigint) (int64, error) {
	res := int64(a + b)
	if a > 0 && b > 0 && res <= 0 {
		return 0, errors.New("result of multiplication is out of the range of type int64")
	}

	return res, nil
}

// Функция вычитания двух больших чисел
func sub(a, b bigint) int64 {
	return int64(a - b)
}

func main() {
	var a bigint = 1<<30 - 233
	var b bigint = 1<<23 - 1001

	if err := checkBigintValue(a); err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	if err := checkBigintValue(b); err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	m, err := mult(a, b)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Printf("%d * %d = %d\n", a, b, m)

	fmt.Printf("%d / %d = %d\n", a, b, div(a, b))

	s, err := sum(a, b)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}
	fmt.Printf("%d + %d = %d\n", a, b, s)

	fmt.Printf("%d - %d = %d\n", a, b, sub(a, b))
}
