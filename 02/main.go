package main

import "fmt"

func main() {
	var a int
	fmt.Print("Введите число: ")
	fmt.Scan(&a)

	if a%2 == 0 {
		fmt.Printf("Число  %v: четное, ", a)
	} else {
		fmt.Printf("Число  %v: нечетное, ", a)
	}

	if a >= 0 {
		fmt.Print("положительное, ")
	} else {
		fmt.Print("отрицательное, ")
	}

	if (a%3 == 0) && (a%5 == 0) {
		fmt.Print("делится на 3 и на 5")
	} else {
		fmt.Print("не делится на 3 и на 5")
	}
}
