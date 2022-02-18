package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Количество горутин, одновременно записывающих данные в map
	// Если при запуске программы флаг не указан, то значение по умолчанию - 10
	workers := flag.Int("workers", 10, "Number of goroutines that will write data to map")
	flag.Parse()

	// С помощью данной структуры реализована мапа, которая является потокобезопасной,
	// то есть, в отличие от первой версии программы, здесь не нужно использовать мьютексы, так как
	// блокировка ресурса происходит в функциях структуры
	m := sync.Map{}

	// Контекст, который автоматически отменяется по истечении трех секунд
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Группа горутин, завершения которых программа будет ожидать
	wg := sync.WaitGroup{}
	wg.Add(*workers)

	// Запускаем горутины в количестве *workers штук
	for i := 0; i < *workers; i++ {
		go func(num int) {
			// По окончанию работы горутины счетчик запущенных горутин в структуре wg уменьшается на единицу
			defer wg.Done()
			count := 0
			for {
				select {
				// Когда пройдет три секунды, канал ctx.Done() окажется закрыт,
				// тогда все горутины прекратят свое выполнение
				case <-ctx.Done():
					return
				default:
					count++
					// Запись данных в структуру
					m.Store(fmt.Sprintf("Goroutine №%d", num), count)
				}
			}
		}(i + 1)
	}

	// Ожидаем выполнения всех горутин
	wg.Wait()

	fmt.Println("How many times each goroutine has accessed the map:")
	for i := 0; i < *workers; i++ {
		key := fmt.Sprintf("Goroutine №%d", i+1)
		// Чтение данных по ключу key
		val, _ := m.Load(key)
		fmt.Printf("%s: %d\n", key, val)
	}
}
