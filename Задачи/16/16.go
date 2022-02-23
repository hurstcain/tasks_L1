package main

import "fmt"

// Quicksort - функция быстрой сортировки
func Quicksort(arr []int, low, high int) {
	if low < high {
		// Индекс опорного элемента, который разделяет массив arr на две части,
		// а затем для каждой из частей массива рекурсивно вызывается функция Quicksort
		p := partition(arr, low, high)
		Quicksort(arr, low, p-1)
		Quicksort(arr, p+1, high)
	}
}

// Функция разбиения массива и получения индекса опорного элемента
func partition(arr []int, low, high int) int {
	// Опорный элемент, всегда является последним элементом части массива
	p := arr[high]
	// Начальный индекс части массива, является возвращаемым значением в качестве индекса опорного элемента,
	// который впоследствии разделит массив arr или его часть на две части
	i := low

	// В цикле все элементы меньше опорного сдвигаются влево
	for j := low; j < high; j++ {
		if arr[j] < p {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// Опорный элемент также сдвигается левее
	arr[i], arr[high] = arr[high], arr[i]

	return i
}

func main() {
	arr := []int{1, 3, 0, 10, 9, 2}
	fmt.Printf("%v -> ", arr)
	Quicksort(arr, 0, len(arr)-1)
	fmt.Printf("%v\n", arr)
}
