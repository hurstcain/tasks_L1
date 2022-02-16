package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	// Группа горутин, завершения которых нужно будет ждать
	var wg sync.WaitGroup
	// Добавляем в группу 5 горутин (столько же элементов в массиве)
	wg.Add(5)

	for _, val := range arr {
		// Горутина выводит на экран квадрат одного из элементов массива arr
		go func(val int) {
			// Вызываем метод wg.Done() после завершения работы горутины, чтобы уменьшить счетчик
			// работающих горутин на единицу
			defer wg.Done()
			fmt.Println(val * val)
		}(val)
	}

	// Ждем завершения работы всех горутин
	wg.Wait()
}
