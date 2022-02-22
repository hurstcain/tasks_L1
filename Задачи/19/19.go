package main

import "fmt"

// ReverseString - переворачивает символы в строке
func ReverseString(s string) string {
	// Слайс символов (рун) строки s
	c := []rune(s)
	// Длина слайса рун
	cLen := len(c)

	// Проходимся по слайсу c и меняем местами противоположные элементы, пока не дойдем до середины слайса
	for i := 0; i < cLen/2; i++ {
		c[i], c[cLen-i-1] = c[cLen-i-1], c[i]
	}

	// Преобразомываем слайс рун в строку и возвращаем ее
	return string(c)
}

func main() {
	s := "главрыба"

	fmt.Printf("%s -> %s\n", s, ReverseString(s))
}
