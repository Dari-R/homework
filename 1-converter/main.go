package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var money float64
	var firstVal, secondVal, repeatVar string
	currencies := map[string]float64{
		"USD": 1,
		"RUB": 79.74,
		"EUR": 0.88,
	}
	fmt.Println("__Калькулятор валют: EUR, USD, RUB__")
	for {
		firstVal, money, secondVal = inputFunc(currencies)
		fmt.Println(converter(firstVal, secondVal, money,currencies))
		fmt.Println("Хотите повторить? (y/n)")
		fmt.Scan(&repeatVar)
		if strings.ToLower(repeatVar) != "y" && strings.ToLower(repeatVar) != "Y" {
			break
		}
	}
}

func inputFunc(currencies map[string]float64) (string, float64, string) {
	var money float64
	var firstVal, secondVal, checkMoney string 
	var err error
	valConv(currencies, &firstVal)
	fmt.Println("Отлично! Введите сумму:")
	for {
		fmt.Scan(&checkMoney)
		money, err = strconv.ParseFloat(checkMoney, 64)
		if err != nil {
			fmt.Println("Неккоректный ввод, введите сумму еще раз")
		} else {
			break
		}
	}
	valConv(currencies, &secondVal)
	return firstVal, money, secondVal
}

func valConv(currencies map[string]float64, val *string){
	fmt.Println("Введите валюту(EUR, USD, RUB)")
	for {
		_, err := fmt.Scan(val)
		*val = strings.ToUpper(*val)
		if err != nil {
			fmt.Println("Ошибка ввода")
		} else if _, ok := currencies[*val]; !ok{
			fmt.Println("Неверная валюта. Доступны только: EUR, USD, RUB. Повторите ввод.")
		} else {
			break
		}
	}
}
func converter(firstVal, secondVal string, money float64, currencies map[string]float64) string {
	var result string
	sum := money*currencies[secondVal]/currencies[firstVal]
	result = fmt.Sprintf("%.2f %s это %.2f в %s \n", money, firstVal, sum, secondVal )
	return result
}
