package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var operation int
	var str string

	for {
		fmt.Println("Что вы хотите вычислить? (1.AVG - среднее 2. SUM - сумму 3. MED - медиану 4.Выход)\n Введите номер:")
		fmt.Scan(&operation)
		if operation != 1 && operation != 2 && operation != 3 && operation != 4 {
			fmt.Println("Неверная операция.")
			continue
		} else if operation == 4 {
			break
		}
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Введите числа через запятую: ")
		str, _ = reader.ReadString('\n')
		resPrint(&operation, str)
		if operation == 4 {
			break
		}
	}
}

func resPrint(operation *int, str string) {
	var numArr []float64
	var err error
	if numArr, err = convStr(str); err != nil {
		fmt.Println(err)
		return
	}
	switch *operation {
	case 1:
		fmt.Println("Среднее:", avg(numArr))
	case 2:
		fmt.Println("Сумма:", sum(numArr))
	case 3:
		fmt.Println("Медиана:", med(numArr))
	}
	fmt.Println("Продолжить? 0 - да, 4 - выход")
	fmt.Scan(operation)
}
func convStr(str string) ([]float64, error) {
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, " ", "")
	parts := strings.Split(str, ",")

	numArr := make([]float64, len(parts))
	for i, val := range parts {
		num, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return nil, fmt.Errorf("ошибка преобразования: %v", err)
		}
		numArr[i] = num
	}
	return numArr, nil
}
func sum(numArr []float64) float64 {
	var sum float64
	for i := 0; i < len(numArr); i++ {
		sum += numArr[i]
	}
	return sum
}

func avg(numArr []float64) float64 {
	avg := sum(numArr) / float64(len(numArr))
	return avg
}

func med(numArr []float64) float64 {
	var num, mediana float64
	if len(numArr) % 2 == 0{
		mediana = avg(numArr)
		return mediana
	}
	for i := 0; i < len(numArr); i++ {
		for j := 0; j < len(numArr)-i-1; j++ {
			if numArr[j+1] < numArr[j] {
				num = numArr[j]
				numArr[j] = numArr[j+1]
				numArr[j+1] = num
			}
		}
	}
	
	return numArr[len(numArr)/2]
}
