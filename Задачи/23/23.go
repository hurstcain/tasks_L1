package main

import (
	"errors"
	"fmt"
	"log"
)

// DeleteElement - удаляет i-ый элемент из слайса arr
func DeleteElement(arr *[]int, i int) error {
	if i >= len(*arr) || i < 0 {
		return errors.New("index value is larger than size of array")
	}
	*arr = append((*arr)[:i], (*arr)[i+1:]...)
	return nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("%v -> ", arr)
	if err := DeleteElement(&arr, 3); err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("%v\n", arr)
	// output: [1 2 3 4 5 6] -> [1 2 3 5 6]
}
