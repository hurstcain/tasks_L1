package main

import (
	"context"
	"fmt"
	"time"
)

// Реализует функцию time.Sleep с помощью функции time.After
func sleep1(seconds int) {
	// Ждет, пока не прошло seconds секунд. Затем возвращает канал со значением текущего времени,
	// данные из канала принимаются, и функция завершает работу
	<-time.After(time.Duration(seconds) * time.Second)
}

// Функция реализует функцию time.Sleep с помощью контекста
func sleep2(seconds int) {
	// Контекст, который отменяется через seconds секунд
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(seconds)*time.Second)
	defer cancel()
	// По истечении seconds секунд канал ctx.Done() закрывается, функция прекращает ожидание и завершает свою работу
	<-ctx.Done()
}

func main() {
	fmt.Printf("Current time: %v\n", time.Now().Format(time.RFC850))

	sleep1(3)

	fmt.Printf("Current time: %v\n", time.Now().Format(time.RFC850))

	sleep2(3)

	fmt.Printf("Current time: %v\n", time.Now().Format(time.RFC850))
}
