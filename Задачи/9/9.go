package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

func main() {
	// Канал, в который записываются числа x
	chanx := make(chan int, 10)
	// Канал, в который записываются квадраты числа x
	chanxx := make(chan int, 10)

	// Канал, который будет принимать сигнал завершения работы программы (Ctrl+C)
	chExit := make(chan os.Signal, 1)
	signal.Notify(chExit, os.Interrupt)

	// Группа горутин, завершения которых будет ожидать программа
	wg := sync.WaitGroup{}

	wg.Add(1)
	// Первая горутина, которая записывает значения x в канал chanx
	// Когда происходит нажатие Ctrl+C, канал chanx закрывается, и горутина заканчивает свою работу
	go func() {
		defer wg.Done()
		for {
			select {
			case <-chExit:
				fmt.Println("Ctrl+C pressed in terminal. Closing all channels...")
				close(chanx)
				return
			default:
				chanx <- rand.Intn(10)
			}
		}
	}()

	wg.Add(1)
	// Горутина, которая читает значения x из канала chanx, а затем записывает квадрат x в канал chanxx
	// Когда все значения из канала chanx прочитаны и канал оказывается закрыт, функция выходит из цикла, закрывает
	// канал chanxx и заканчивает свою работу
	go func() {
		defer wg.Done()
		for x := range chanx {
			chanxx <- x * x
		}
		close(chanxx)
	}()

	wg.Add(1)
	// Горутина, которая читает квадраты x из канала chanxx и выводит эти значения на экран
	// Когда все значения прочитаны и канал chanxx оказывается закрыт, горутина прекращает свою работу
	go func() {
		defer wg.Done()
		for val := range chanxx {
			fmt.Println(val)
		}
	}()

	// Ожидание завершения работы всех трех горутин
	wg.Wait()

	fmt.Println("Program is finished")
}
