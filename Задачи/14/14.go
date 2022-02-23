package main

import (
	"errors"
	"fmt"
	"reflect"
)

// Определение типа переменной a с помощью возможностей форматирования строки
func defineType1(a interface{}) string {
	return fmt.Sprintf("%T", a)
}

// Определение типа переменной a с помощью функции TypeOf() из библиотеки reflect
func defineType2(a interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(a))
}

// Определение типа переменной a с помощью оператора switch.
// Функция возвращает пустую строку и ошибку в случае, когда тип a не является одним из трех типов: int, string, bool
func defineType3(a interface{}) (string, error) {
	var aType string

	switch a.(type) {
	case int:
		aType = "int"
	case string:
		aType = "string"
	case bool:
		aType = "bool"
	default:
		return "", errors.New("failed to determine the type of variable")
	}

	return aType, nil
}

func main() {
	var a interface{}

	a = 1
	fmt.Printf("a = %v:\n", a)
	fmt.Printf("defineType1(a) -> %s\n", defineType1(a))
	fmt.Printf("defineType2(a) -> %s\n", defineType2(a))
	aType, _ := defineType3(a)
	fmt.Printf("defineType3(a) -> %s\n\n", aType)
	// output:
	// a = 1:
	// defineType1(a) -> int
	// defineType2(a) -> int
	// defineType3(a) -> int

	a = "str"
	fmt.Printf("a = %#v:\n", a)
	fmt.Printf("defineType1(a) -> %s\n", defineType1(a))
	fmt.Printf("defineType2(a) -> %s\n", defineType2(a))
	aType, _ = defineType3(a)
	fmt.Printf("defineType3(a) -> %s\n\n", aType)
	// output:
	// a = "str":
	// defineType1(a) -> string
	// defineType2(a) -> string
	// defineType3(a) -> string

	a = true
	fmt.Printf("a = %v:\n", a)
	fmt.Printf("defineType1(a) -> %s\n", defineType1(a))
	fmt.Printf("defineType2(a) -> %s\n", defineType2(a))
	aType, _ = defineType3(a)
	fmt.Printf("defineType3(a) -> %s\n\n", aType)
	// output:
	// a = true:
	// defineType1(a) -> bool
	// defineType2(a) -> bool
	// defineType3(a) -> bool

	ch := make(chan interface{}, 1)

	ch <- 1
	fmt.Println("ch <- 1:")
	fmt.Printf("defineType1(<-ch) -> %s\n", defineType1(<-ch))
	ch <- "str"
	fmt.Println("ch <- \"str\":")
	fmt.Printf("defineType2(<-ch) -> %s\n", defineType2(<-ch))
	ch <- false
	fmt.Println("ch <- false:")
	chType, _ := defineType3(<-ch)
	fmt.Printf("defineType3(<-ch) -> %s\n", chType)
	// output:
	// ch <- 1:
	// defineType1(<-ch) -> int
	// ch <- "str":
	// defineType2(<-ch) -> string
	// ch <- false:
	// defineType3(<-ch) -> bool

}
