package main

import "fmt"

func printTable(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%v x %v = %v\n", i, n, i*n)
	}
}

func main() {
	var n int
	fmt.Print("Введите число от 1 до 10: ")
	fmt.Scan(&n)

	fmt.Printf("Таблица умножения на %v:\n", n)
	printTable(n)

}
