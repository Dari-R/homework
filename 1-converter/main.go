package main

import (
	"fmt"
	"strconv"
	"strings"
)

const usdTorub = 79.74
const usdToeur = 0.88
const usd, eur, rub = "USD", "EUR", "RUB"

func inputFunc() (string, float64, string) {
	var money float64
	var firstVal, secondVal string
	fmt.Println("Введите валюту которую хотите конвертировать(EUR, USD, RYB)")
	for {
		_, err := fmt.Scan(&firstVal)
		firstVal = strings.ToUpper(firstVal)
		if err != nil {
			fmt.Println("Ошибка ввода")
		} else if firstVal != usd && firstVal != eur && firstVal != rub {
			fmt.Println("Ошибка, еще раз введите валюту которую хотите конвертировать(EUR, USD, RYB)")
		} else {
			break
		}
	}
	fmt.Println("Отлично! Введите сумму:")
	var checkMoney string
	var err error
	for {
		fmt.Scan(&checkMoney)
		money, err = strconv.ParseFloat(checkMoney, 64)
		if err != nil {
			fmt.Println("Неккоректный ввод, введите сумму еще раз")
		} else {
			break
		}
	}
	fmt.Println("Введите валюту  в которую хотите перевести(EUR, USD, RYB)")
	for {
		_, err := fmt.Scan(&secondVal)
		secondVal = strings.ToUpper(secondVal)
		if err != nil {
			fmt.Println("Ошибка ввода")
		} else if secondVal != usd && secondVal != eur && secondVal != rub {
			fmt.Println("Ошибка, еще раз введите валюту которую хотите конвертировать(EUR, USD, RYB)")
		} else {
			break
		}
	}
	return strings.ToUpper(firstVal), money, strings.ToUpper(secondVal)
}

func converter(firstVal, secondVal string, money float64) string {

	var result string
	switch {
	case firstVal == rub && secondVal == usd:
		result = fmt.Sprintf("%.f рублей это %.2f долларах \n", money, money/usdTorub)
	case firstVal == rub && secondVal == eur:
		result = fmt.Sprintf("%.f рублей это %.2f евро\n", money, money/usdTorub*usdToeur)
	case firstVal == usd && secondVal == rub:
		result = fmt.Sprintf("%.f долларов это %.2f рублей\n", money, money*usdTorub)
	case firstVal == usd && secondVal == eur:
		result = fmt.Sprintf("%.f долларов это %.2f евро\n", money, money*usdToeur)
	case firstVal == eur && secondVal == rub:
		result = fmt.Sprintf("%.f евро это %.2f рублей\n", money, money/usdToeur*usdTorub)
	case firstVal == eur && secondVal == usd:
		result = fmt.Sprintf("%.f евро это %.2f долларов\n", money, money/usdToeur)
	case firstVal == eur && secondVal == eur:
		result = fmt.Sprintf("%.f евро это %.2f евро\n", money, money)
	case firstVal == usd && secondVal == usd:
		result = fmt.Sprintf("%.f долларов это %.2f долларов\n", money, money)
	case firstVal == rub && secondVal == rub:
		result = fmt.Sprintf("%.f рублей это %.2f рублей\n", money, money)
	}
	return result
}
func main() {
	fmt.Println("__Калькулятор для перевода из евро в рубли__")
	var money float64
	var firstVal, secondVal, repeatVar string
	for {
		firstVal, money, secondVal = inputFunc()
		fmt.Println(converter(firstVal, secondVal, money))
		fmt.Println("Хотите повторить? (y/n)")
		fmt.Scan(&repeatVar);
		if repeatVar != "y" && repeatVar != "Y"{
			break
		}
	}
}
