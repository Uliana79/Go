package main

import "fmt"

func main() {
	var a, b float64
	var op string
	fmt.Print("Введите два числа: ")
	fmt.Scan(&a, &b)
	fmt.Print("Введите операцию (+, -, * или /): ")
	fmt.Scan(&op)

	switch op {
	case "+":
		fmt.Println("Результат: ", a+b)
	case "-":
		fmt.Println("Результат: ", a-b)
	case "*":
		fmt.Println("Результат: ", a*b)
	case "/":
		if b == 0 {
			fmt.Println("На ноль делить нельзя!")
		} else {
			fmt.Println("Результат: ", a/b)
		}
	}

}
