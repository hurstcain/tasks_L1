package main

import (
	"fmt"
	"strings"
)

// ReverseWords - переворачивает слова в строке s
func ReverseWords(s string) string {
	// Слайс со словами из строки s
	arr := strings.Split(s, " ")
	// Длина слайса arr
	arrLen := len(arr)

	// Проходимся по слайсу и меняем противоположные элементы местами, пока не дойдем до середины слайса
	for i := 0; i < arrLen/2; i++ {
		arr[i], arr[arrLen-i-1] = arr[arrLen-i-1], arr[i]
	}

	return strings.Join(arr, " ")
}

func main() {
	s := "snow dog sun"

	fmt.Printf("%s -> %s", s, ReverseWords(s))
	// output: snow dog sun -> sun dog snow
}
