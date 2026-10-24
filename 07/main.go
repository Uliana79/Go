package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	sum := 0

	for i := 1; i <= n; i++ {
		if (i%2 == 0) && (i%3 == 0) {
			sum += i
		}
	}

	fmt.Printf("Сумма чисел (удовлетворяющих двум условиям) от 1 до %v : %v\n", n, sum)

}
