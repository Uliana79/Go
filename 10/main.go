package main

import "fmt"

func main() {
	var n, even, odd, a, count, max, min int
	even, odd, count, max, min = 0, 0, 0, 0, 0
	var sred float64

	fmt.Print("Введите количество чисел: ")
	fmt.Scan(&n)
	fmt.Print("Введите числа: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		count += a

		if a > max {
			max = a
		}
		if a < min {
			min = a
		}

		if a%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	sred = float64(count) / float64(n)

	fmt.Println("Сумма всех введённых чисел: ", count)
	fmt.Println("Среднее значение: ", sred)
	fmt.Println("Максимальное значение: ", max)
	fmt.Println("Минимальное значение: ", min)
	fmt.Println("Количество чётных чисел: ", even)
	fmt.Println("Количество нечётных чисел: ", odd)

}
