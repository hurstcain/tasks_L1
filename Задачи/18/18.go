package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

// Counter - структура счетчик.
type Counter struct {
	// Значение счетчика
	count int
	// Встраиваем структуру sync.Mutex, чтобы пользоваться ее методами
	sync.Mutex
}

// Increment - увеличивает значение счетчика на единицу.
func (c *Counter) Increment() {
	c.Lock()
	c.count++
	c.Unlock()
}

// Count - возвращает значение счетчика.
func (c *Counter) Count() int {
	c.Lock()
	defer c.Unlock()
	return c.count
}

func main() {
	c := Counter{}

	// Количество горутин, которые будут инкрементировать счетчик
	const goroutinesCount = 13
	// Группа горутин, завершения которых программа будет ожидать
	wg := sync.WaitGroup{}
	// Добавляем в эту группу goroutinesCount горутин
	wg.Add(goroutinesCount)

	// Канал, принимающий сигнал завершения работы программы (нажатие Ctrl+C)
	chExitSignal := make(chan os.Signal, 1)
	signal.Notify(chExitSignal, os.Interrupt)
	// Канал, который будет закрыт после нажатия Ctrl+C, что оповестит все остальные горутины о том,
	// что пора завершать работу
	chExit := make(chan struct{})

	// Горутина, ожидающая записи в канал chExitSignal. Когда в этот канал записываются данные, горутина
	// закрывает канал chExit и завершает работу
	go func() {
		select {
		case <-chExitSignal:
			close(chExit)
		}
	}()

	// Запуск горутин, инкрементирующих значение счетчика. Горутины завершают свою работу,
	// когда канал chExit оказывается закрыт
	for i := 0; i < goroutinesCount; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-chExit:
					return
				default:
					c.Increment()
				}
			}
		}()
	}

	// Ожидание завершения работы горутин
	wg.Wait()

	fmt.Printf("Значение счетчика: %d", c.Count())
	// output: Значение счетчика: 238524928
}
