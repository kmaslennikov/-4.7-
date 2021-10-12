package main

import "fmt"

func main() {
	amountToRate30 := 50000
	amountToRate20 := 10000

	var income int   //доход
	var totalTax int //сумма налога

	var totalIncome int //значение для вывода сообщения

	fmt.Println("Введите сумму дохода:")
	fmt.Scan(&income)

	totalIncome = income

	if income > amountToRate30 {
		more50 := (income - amountToRate30) // сумма больше 50к
		totalTax += (more50 / 100) * 30
		//остаток для расчета оставшейся суммы налога
		income = amountToRate30
	}

	if income > amountToRate20 {
		more10 := income - amountToRate20
		totalTax += (more10 / 100) * 20

		income = amountToRate20
	}

	totalTax += (income / 100) * 13

	fmt.Println("С дохода", totalIncome, "рублей придётся заплатить:", totalTax)

}
