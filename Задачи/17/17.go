package main

import "fmt"

// BinarySearch - функция бинарного поиска элемента el в массиве arr.
// Если элемент найден, то функция возвращает индекс этого элемента в массиве,
// а если элемент не найден, то функция возвращает -1.
func BinarySearch(arr []int, el int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if el < arr[mid] {
			right = mid - 1
		} else if el > arr[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	arr := []int{0, 1, 20, 30, 34, 25}
	el := 34
	if i := BinarySearch(arr, el); i != -1 {
		fmt.Printf("Элемент %d найден. Его индекс в массиве: %d\n", el, i)
	} else {
		fmt.Println("Элемент не найден")
	}
	// output: Элемент 34 найден. Его индекс в массиве: 4
}
