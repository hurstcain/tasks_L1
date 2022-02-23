package main

/*
Программа демонстрирует остановку выполнения горутины с помощью context.WithDeadline()
*/

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Контекст, который завершается через 6 секунд после запуска программы
	// (Ко времени, когда программа начала работать, добавляем 6 секунд, получается дедлайн -
	// время запуска программы + 6 секунд)
	cxt, cancel := context.WithDeadline(context.Background(), time.Now().Add(6*time.Second))
	defer cancel()

	// Группа горутин, завершения которых будет ожидать программа
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Горутина, которая каждые 100 миллисекунд выводит рандомное число от 0 до 9.
	// Если канал cxt.Done() оказывается закрыт (текущее время = время дедлайна),
	// то горутина завершает работу.
	go func() {
		defer wg.Done()
		for {
			select {
			case <-cxt.Done():
				return
			default:
				fmt.Println(rand.Intn(10))
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	// Ожидание завершения работы горутины
	wg.Wait()

	fmt.Println("Программа завершила работу")
}
