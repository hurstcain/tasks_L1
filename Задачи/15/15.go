package main

import (
	"math/rand"
)

var justString string

func createHugeString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}

// Данная функция создает строку v длиной 1024 символа и присваивает переменной justString слайс из первых ста
// символов строки v. Строка в Go является слайсом байт или рун. В данном случае происходит так,
// что строка (слайс рун) justString будет ссылаться в памяти на массив, на который ссылается сама строка v. По идее,
// переменная v является временной переменной, она существует в памяти, пока выполняется функция someFunc().
// Но, если на значение этой переменной будет ссылаться глобальная переменная (которой и является justString),
// то, после завершения функции, переменная v не будет удалена из памяти. Таким образом, явно присваивая слайс рун,
// созданный из слайса v, переменной justString, получается, что в памяти также сохранится большая строка v,
// а это не очень хорошо, так как по идее такая большая строка нам не нужна, занимается лишняя память.
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

// Чтобы исправить это, нужно скопировать часть данных с помощью функции copy(). Тогда justString не будет ссылаться
// на строку v, и после выхода из функции память будет очищена от ненужных данных.
func someFunc2() {
	v := createHugeString(1 << 10)
	// Создаем временный слайс рун длиной и емкостью 100
	temp := make([]rune, 100)
	// Копируем в слайс первые 100 символов строки v
	copy(temp, []rune(v[:100]))
	// Преобразовываем слайс рун в строку и присваиваем ее justString. Таким образом, по выходу из функции
	// переменные v и temp удалятся из памяти, останется только justString
	justString = string(temp)
}
