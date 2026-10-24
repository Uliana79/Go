package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	var sred float64
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	sred = float64(sum) / float64(n)
	fmt.Printf("Сумма всех чисел от 1 до %v: %v\n", n, sum)
	fmt.Println("Среднее значение: ", sred)

}
