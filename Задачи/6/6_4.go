package main

/*
Программа демонстрирует остановку выполнения горутины с помощью context.WithTimeout()
*/

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Контекст, который завершается по таймауту, то есть через 4 секунды
	cxt, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	// Группа горутин, завершения которых будет ожидать программа
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Горутина, которая каждые 100 миллисекунд выводит рандомное число от 0 до 9.
	// Если канал cxt.Done() оказывается закрыт (прошло 4 секунды),
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
