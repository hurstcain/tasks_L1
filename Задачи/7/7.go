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

	// Map, куда будут записываться данные
	m := make(map[string]int)
	// Мьютекс, блокирующий доступ к мапе m
	mu := sync.Mutex{}

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
			for {
				select {
				// Когда пройдет три секунды, канал ctx.Done() окажется закрыт,
				// тогда все горутины прекратят свое выполнение
				case <-ctx.Done():
					return
				default:
					mu.Lock()
					// Ключ в мапе - информация о горутине (ее номер)
					// Значение - сколько раз данная горутина получила доступ на запись в мапу
					// Каждый раз, когда какая-либо горутина получает доступ на запись, она увеличивает
					// значение по ключу с ее номером на единицу
					m[fmt.Sprintf("Goroutine №%d", num)]++
					mu.Unlock()
				}
			}
		}(i + 1)
	}

	// Ожидаем выполнения всех горутин
	wg.Wait()

	fmt.Println("How many times each goroutine has accessed the map:")
	for key, val := range m {
		fmt.Printf("%s: %d\n", key, val)
	}
	// output:
	// Goroutine №2: 1574626
	// Goroutine №6: 1600554
	// Goroutine №5: 1467768
	// Goroutine №7: 1484413
	// Goroutine №3: 1572974
	// Goroutine №4: 1614989
	// Goroutine №1: 1490189
	// Goroutine №10: 1575373
	// Goroutine №8: 1531322
	// Goroutine №9: 1551141
}
