package internal

import "fmt"

func PrintReceipt(amount int, commission int, total int) {
	fmt.Println("\n================================")
	fmt.Println("        ЭЛЕКТРОННЫЙ ЧЕК        ")
	fmt.Println("           alif mobi           ")
	fmt.Println("================================")
	fmt.Printf(" Сумма перевода:  %10d сум\n", amount)
	fmt.Printf(" Комиссия:        %10d сум\n", commission)
	fmt.Println("--------------------------------")
	fmt.Printf(" Итого к оплате:  %10d сум\n", total)
	fmt.Println("================================")
	fmt.Println("     Спасибо за перевод!       ")
	fmt.Println("================================")
}
