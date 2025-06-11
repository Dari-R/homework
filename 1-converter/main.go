package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var currencies = map[string]float64{
	"USD": 1.0,
	"RUB": 79.74,
	"EUR": 0.88,
}

var operations = map[string]func(float64, string, string) string{
	"CONVERT": func(amount float64, from, to string) string {
		fromRate, fromOk := currencies[strings.ToUpper(from)]
		toRate, toOk := currencies[strings.ToUpper(to)]
		if !fromOk || !toOk {
			return "Ошибка: Неверная валюта. Доступны только EUR, USD, RUB."
		}
		converted := amount * toRate / fromRate
		return fmt.Sprintf("%.2f %s это %.2f в %s", amount, strings.ToUpper(from), converted, strings.ToUpper(to))
	},
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Использование: go run main.go <из_валюты> <сумма> <в_валюту>")
		os.Exit(1)
	}

	from := os.Args[1]
	amountStr := os.Args[2]
	to := os.Args[3]

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Ошибка: сумма должна быть числом")
		os.Exit(1)
	}

	op := "CONVERT"
	if operation, ok := operations[op]; ok {
		fmt.Println(operation(amount, from, to))
	} else {
		fmt.Println("Ошибка: операция не поддерживается")
	}
}
