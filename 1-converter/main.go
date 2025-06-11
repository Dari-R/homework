package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Использование: go run main.go <FROM> <TO> <AMOUNT>")
		return
	}

	from := os.Args[1]
	to := os.Args[2]
	amountStr := os.Args[3]

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Ошибка: некорректная сумма.")
		return
	}

	currencies := map[string]float64{
		"USD": 1.0,
		"EUR": 0.93,
		"JPY": 151.67,
		"GBP": 0.79,
		"CNY": 7.25,
	}

	result, err := convert(from, to, amount, &currencies)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("%.2f %s = %.2f %s\n", amount, from, result, to)
}

func convert(from, to string, amount float64, rates *map[string]float64) (float64, error) {
	fromRate, ok := (*rates)[from]
	if !ok {
		return 0, fmt.Errorf("неизвестная валюта: %s", from)
	}
	toRate, ok := (*rates)[to]
	if !ok {
		return 0, fmt.Errorf("неизвестная валюта: %s", to)
	}
	usdAmount := amount / fromRate
	return usdAmount * toRate, nil
}
