package main

import (
	"fmt"
	"unicode"
)

// Проверяет, все ли символы в строке уникальные. Если да, то возвращает true, а если нет, то false.
// Функция не чувствительна к регистру
func areSymbolsUnique(s string) bool {
	// Мапа, где в качестве ключей хранятся unicode символы из строки s, а в качестве значений хранятся
	// пустые структуры, так как значения мапы нам не особо нужны
	symbolsMap := make(map[rune]struct{})

	// Цикл проходится по каждой руне в строке s
	for _, r := range s {
		// Приводим символ к нижнему регистру
		lowRune := unicode.ToLower(r)
		// Если до этого такой же символ был записан в качестве ключа в мапу symbolsMap, то функция возвращает false
		if _, ok := symbolsMap[lowRune]; ok {
			return false
		}
		// Если символ до этого не встречался, то добавляем ключ с этим символом в мапу
		symbolsMap[lowRune] = struct{}{}
	}

	// Если цикл дошел до конца и повторяющихся символов в строке не было найдено, то функция возвращает true
	return true
}

func main() {
	fmt.Printf("areSymbolsUnique(\"abcd\") -> %v\n", areSymbolsUnique("abcd"))
	fmt.Printf("areSymbolsUnique(\"абвгд\") -> %v\n", areSymbolsUnique("абвгд"))
	fmt.Printf("areSymbolsUnique(\"абвгдА\") -> %v\n", areSymbolsUnique("абвгдА"))
	fmt.Printf("areSymbolsUnique(\"abCdefAaf\") -> %v\n", areSymbolsUnique("abCdefAaf"))
	fmt.Printf("areSymbolsUnique(\"abBA\") -> %v\n", areSymbolsUnique("abBA"))
	fmt.Printf("areSymbolsUnique(\"aabcd\") -> %v\n", areSymbolsUnique("aabcd"))
}
