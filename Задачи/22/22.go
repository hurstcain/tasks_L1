// Для реализации больших чисел и операций над ними
// используется библиотека math/big
package main

import (
	"fmt"
	"math/big"
)

// BigInt - структура, описывающая большое число.
type BigInt struct {
	a *big.Int
}

// NewBigInt - создает и возвращает экземпляр структуры BigInt
func NewBigInt() BigInt {
	return BigInt{a: new(big.Int)}
}

// Get - возвращает значение большого числа
func (bi *BigInt) Get() *big.Int {
	return bi.a
}

// Set - устанавливает значение большого числа
func (bi *BigInt) Set(s string) {
	bi.a.SetString(s, 10)
}

// Mul - функция умножения больших чисел
func Mul(a, b BigInt) BigInt {
	res := NewBigInt()
	res.Get().Mul(a.Get(), b.Get())

	return res
}

// Div - функция деления больших чисел
func Div(a, b BigInt) BigInt {
	res := NewBigInt()
	res.Get().Div(a.Get(), b.Get())

	return res
}

// Sum -функция суммирования больших чисел
func Sum(a, b BigInt) BigInt {
	res := NewBigInt()
	res.Get().Add(a.Get(), b.Get())

	return res
}

// Sub - функция вычитания больших чисел
func Sub(a, b BigInt) BigInt {
	res := NewBigInt()
	res.Get().Sub(a.Get(), b.Get())

	return res
}

func main() {
	a := NewBigInt()
	b := NewBigInt()
	a.Set("12390909120982392340921830178401290812471070")
	b.Set("827182980001209820012983718203981028")

	mul := Mul(a, b)
	fmt.Printf("%v * %v = %v\n", a.Get(), b.Get(), mul.Get())
	// output: 12390909120982392340921830178401290812471070 * 827182980001209820012983718203981028 = 10249549131618386593926522424243779064165327958720157188084233952493064078859960

	div := Div(a, b)
	fmt.Printf("%v / %v = %v\n", a.Get(), b.Get(), div.Get())
	// output: 12390909120982392340921830178401290812471070 / 827182980001209820012983718203981028 = 14979647

	sum := Sum(a, b)
	fmt.Printf("%v + %v = %v\n", a.Get(), b.Get(), sum.Get())
	// output: 12390909120982392340921830178401290812471070 + 827182980001209820012983718203981028 = 12390909948165372342131650191385009016452098

	sub := Sub(a, b)
	fmt.Printf("%v - %v = %v\n", a.Get(), b.Get(), sub.Get())
	// output: 12390909120982392340921830178401290812471070 - 827182980001209820012983718203981028 = 12390908293799412339712010165417572608490042
}
