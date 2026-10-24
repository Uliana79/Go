package main

import "fmt"

func main() {
	var n, plus, minus, zero int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)
	var a float64
	plus, minus, zero = 0, 0, 0

	fmt.Print("Введите числа: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a > 0 {
			plus++
		} else if a < 0 {
			minus++
		} else {
			zero++
		}
	}

	fmt.Println("Количество положительных чисел:", plus)
	fmt.Println("Количество отрицательных чисел:", minus)
	fmt.Println("Количество нулей:", zero)

}
