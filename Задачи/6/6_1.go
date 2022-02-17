package main

/*
Программа демонстрирует остановку работы горутины с помощью закрытия канала stopCh
*/

import (
	"fmt"
	"sync"
	"time"
)

// Выводит на экран значение переменной d каждую секунду
func goroutine(stopCh chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	d := 1
	for {
		select {
		// Если канал stopCh закрыли или записали в него данные, то горутина заканчивает свою работу
		case <-stopCh:
			fmt.Println("Прошо 10 секунд. Завершение работы горутины...")
			return
		default:
			fmt.Println(d)
			d++
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	// Канал, использующийся для того, чтобы просигнализировать горутине о том, когда следует завершать свою работу
	stopCh := make(chan interface{})

	// Группа горутин, завершения которых программа будет ожидать
	// Всего в группе будет одна горутина
	wg := sync.WaitGroup{}

	wg.Add(1)
	// Запускаем горутину
	go goroutine(stopCh, &wg)

	// Ждем 10 секунд
	time.Sleep(10 * time.Second)
	// Закрываем канал, чтобы оповестить горутину о том, что пора заканчивать работу
	close(stopCh)

	// Ждем завершения работы горутины
	wg.Wait()

	fmt.Println("Программа завершила работу")
}
