package main

import "fmt"

func main() {
	var balance float64
	var op int
	fmt.Print("Введите баланс: ")
	fmt.Scan(&balance)
	isExit := false

	fmt.Println("\n1. Посмотреть баланс.\n2. Пополнить баланс.\n3. Снять средства.\n4. Посмотреть меню.\nЧтобы выйти нажмите любую другую кнопку\n")
	for {
		fmt.Print("Выберите одно из действий (1, 2, 3, 4): ")
		fmt.Scan(&op)

		switch op {
		case 1:
			fmt.Println("Баланс: ", balance)
		case 2:
			var a float64
			fmt.Print("Введите сумму на которую пополнить баланс: ")
			fmt.Scan(&a)
			balance += a
			fmt.Println("Баланс обновлен!")
		case 3:
			var a float64
			fmt.Print("Введите сумму которую необходимо снять: ")
			fmt.Scan(&a)
			if balance-a < 0 {
				fmt.Println("Денег на балансе не достаточно!")
			} else {
				balance -= a
				fmt.Println("Баланс обновлен!")
			}
		case 4:
			fmt.Println("\n1. Посмотреть баланс.\n2. Пополнить баланс.\n3. Снять средства.\n4. Посмотреть меню.\nЧтобы выйти нажмите любую другую кнопку\n")
		default:
			fmt.Println("До скорой встречи!")
			isExit = true
		}

		if isExit {
			break
		}
	}
}
