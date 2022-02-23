package main

/*
Программа демонстрирует остановку выполнения горутины с помощью context.WithCancel()
*/

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// WaitGroup
// close
// withCancel cancel()
// time.after
// Timeout

func main() {
	// Контекст, который завершается после вызова cancel()
	cxt, cancel := context.WithCancel(context.Background())

	// Группа горутин, завершения которых будет ожидать программа
	wg := sync.WaitGroup{}
	wg.Add(1)

	// Горутина, которая каждые 100 миллисекунд выводит рандомное число от 0 до 9.
	// Если канал cxt.Done() оказывается закрыт (что означает, что была вызвана функция cancel()),
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

	// Программа ждет 5 секунд
	time.Sleep(5 * time.Second)
	// А потом вызывает функцию
	cancel()

	// Ожидание завершения работы горутины
	wg.Wait()

	fmt.Println("Программа завершила работу")
}
