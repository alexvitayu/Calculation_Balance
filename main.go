package main

import "fmt"

// В цикле спрашиваем ввод транзакций: -10, 10, 40.5
// Добавим каждую в массив транзакций
// Вывести массив
// Вывести сумму баланса в консоль

func main() {
	transactions := []float64{}
	for {
		transaction := inputData()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	summ := calculateBalance(transactions)
	fmt.Println(transactions)
	//fmt.Printf("Ваш баланс составляет: %.1f рублей", summ)
	fmt.Println(summ)
	welcome := greeting()
	fmt.Println(welcome)

}
func inputData() float64 {
	var data float64
	fmt.Println("Введите транзакцию или нажмите n для выхода:")
	fmt.Scan(&data)
	return data
}
func calculateBalance(transactions []float64) float64 {
	balance := 0.0
	for _, value := range transactions {
		balance = balance + value
	}
	return balance
}

func greeting() string {
	var greeting string
	greeting = "Welcome to Belarus"
	//fmt.Println("Welcome to Belarus")
	return greeting
}
