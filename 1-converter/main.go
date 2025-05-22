package main

import "fmt"

func main(){
	const usdToryb = 79.68
	const usdToeur = 0.89
	var money float64
	fmt.Println("__Калькулятор для перевода из евро в рубли__")
	fmt.Println("Введите сумму в евро, которую хотите конвертировать в рубли")
	fmt.Scan(&money)
	fmt.Printf("%.f рублей это %.2f Евро\n", money, money/usdToryb*usdToeur)
}