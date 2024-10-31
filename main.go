package main

import "fmt"

// В цикле спрашиваем ввод транзакций: -10, 10, 40.5
// Добавим каждую в массив транзакций
// Вывести массив
// Вывести сумму баланса в консоль

func main() {
transactions := []float64{}
var data float64
for {
data = inputTransactions()
if data == 0 {
	break
}
transactions = append(transactions, data)
}	
total := totalTransactions(transactions)
fmt.Println(transactions)
fmt.Printf("Сумма ваших транзакций составляет: %.2f", total)
}

func inputTransactions ()  float64 {
	var data float64
fmt.Println("Введите новую транзакцию:")
fmt.Scan(&data)
return data
}

func totalTransactions (transactions []float64) float64 {
	data := 0.0
	for _,value := range transactions {
		data += value
	}
	return data

}
