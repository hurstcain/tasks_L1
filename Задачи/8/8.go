package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
)

// Функция, которая инверсирует i-ый бит числа
// Возвращает новое значение числа с инверсированным битом и ошибку в случае некорректного значения номера бита
// Нумерация битов начинается с конца, первый бит находится в самом конце двоичного представления числа
// Например, в числе 11110 бит со значением 0 является первым битом
func reverseBit(a int64, i int) (int64, error) {
	if i < 1 || i > 63 {
		return 0, errors.New("invalid bit value")
	}

	var aNew int64
	var aTemp int64

	if a < 0 {
		aTemp = -a
	} else {
		aTemp = a
	}

	i--
	if aTemp&(1<<i) == 0 {
		aNew = aTemp | (1 << i)
	} else {
		aNew = aTemp & ^(1 << i)
	}

	if a < 0 {
		return -aNew, nil
	}
	return aNew, nil
}

func main() {
	// Число, бит которого будет инверсирован
	// Если флаг не указан, то значение по умолчанию - 1997
	a := flag.Int64("a", 1997, "Number")
	// Номер бита, по умолчанию значение - 9
	i := flag.Int("bit", 9, "Number's bit to reverse")
	flag.Parse()

	aNew, err := reverseBit(*a, *i)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	fmt.Printf("%d -> %b\n", *a, *a)
	fmt.Printf("%d -> %b\n", aNew, aNew)
	// Output:
	// 1997 -> 11111001101
	// 1741 -> 11011001101
}
