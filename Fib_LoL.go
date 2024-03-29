package main

/* Учебная задача с сайта stepik.org

Дано натуральное число A > 1.
Определите, каким по счету числом Фибоначчи оно является,
то есть выведите такое число n, что φn=A.
Если А не является числом Фибоначчи, выведите число -1.

Входные данные: Вводится натуральное число.

Выходные данные: Выведите ответ на задачу.
*/

import (
	"fmt"
)

func main() {
	var (
		a     int
		fib_1 = 1
		fib_2 = 1
	)
	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fib_2 = fib_1 + fib_2
		fib_1 = fib_2 - fib_1
		if fib_2 == a {
			fmt.Printf("%d ", i+3)
			break
		} else if fib_2 > a {
			fmt.Print(-1)
			break
		}
	}
}
