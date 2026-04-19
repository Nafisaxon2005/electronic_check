package main

import (
	"check/internal"
	Commission "check/internal/commission"
	"fmt"
)

func main() {
	var amount int
	var cardType int

	for {
		fmt.Print("Введите сумму перевода(или 0 для выхода): ")
		fmt.Scan(&amount)

		if amount == 0 {
			fmt.Println("Программа завершена.")
			break
		}

		if amount < 500 || amount > 15000000 {
			fmt.Println("Ошибка: сумма должна быть от 500 до 15 000 000 сум")
			fmt.Println("Пожалуйста, попробуйте еще раз.")
			continue
		}

		fmt.Print("Тип карты (1 - Alif, 2 - Другая): ")
		fmt.Scan(&cardType)

		isAlif := cardType == 1
		var comm, total = Commission.Calculate(amount, isAlif)
		internal.PrintReceipt(amount, comm, total)

		break
	}
}
