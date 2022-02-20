package main

import "fmt"

// Возвращает пересечение двух множеств a и b
func intersection(a []interface{}, b []interface{}) []interface{} {
	// Пересечение множеств
	intersectionSet := make([]interface{}, 0)
	// Мапа, в ключах которой будут храниться элементы множества a
	aMap := make(map[interface{}]struct{})

	for _, val := range a {
		aMap[val] = struct{}{}
	}

	for _, val := range b {
		// Проверяем, если в множестве a есть элемент множества b, то добавляем этот элемент в пересечение
		if _, ok := aMap[val]; ok {
			intersectionSet = append(intersectionSet, val)
		}
	}

	return intersectionSet
}

func main() {
	a := []interface{}{1, 3, 0, "set", 20, "0", 21}
	b := []interface{}{1, "set", "0", 21, 0, 73, 902, "1", "3"}

	intersectionSet := intersection(a, b)

	for _, v := range intersectionSet {
		fmt.Printf("%v (%T); ", v, v)
	}
}
