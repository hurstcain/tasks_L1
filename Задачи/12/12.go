package main

import "fmt"

// Создает множество для последовательности строк a
func createSet(a []string) []string {
	// Множество
	set := make([]string, 0)
	// Мапа, в которой в качестве ключей хранятся значения из слайса a
	aMap := make(map[string]struct{})

	for _, val := range a {
		aMap[val] = struct{}{}
	}

	// Проходимся по ключам aMap и записываем их в множество set. Так как в мапе не может быть нескольких ключей
	// с одинаковым значением, получается, что в set записывается по одному уникальному значению из слайса a
	for key := range aMap {
		set = append(set, key)
	}

	return set
}

func main() {
	a := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Printf("%v -> %v", a, createSet(a))
}
