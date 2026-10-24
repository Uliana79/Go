package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)
	var max, min, a float64
	max, min = 0, 0
	fmt.Print("Введите числа: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		if a > max {
			max = a
		}
		if a < min {
			min = a
		}
	}

	fmt.Println("Максимальное значение: ", max)
	fmt.Println("Минимальное значение: ", min)

}
