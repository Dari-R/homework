package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Пример использования: go run main.go AVG 2,10,9")
		return
	}

	op := strings.ToUpper(os.Args[1])
	rawNumbers := strings.Split(os.Args[2], ",")
	var numbers []float64

	for _, s := range rawNumbers {
		num, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		if err != nil {
			fmt.Printf("Ошибка преобразования числа: %v\n", err)
			return
		}
		numbers = append(numbers, num)
	}

	operations := map[string]func([]float64) float64{
		"SUM": sum,
		"AVG": avg,
		"MED": med,
	}

	fn, ok := operations[op]
	if !ok {
		fmt.Println("Операция не поддерживается. Используйте: SUM, AVG или MED.")
		return
	}

	result := fn(numbers)
	fmt.Printf("Результат %s: %.2f\n", op, result)
}

func sum(nums []float64) float64 {
	var total float64
	for _, n := range nums {
		total += n
	}
	return total
}

func avg(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	return sum(nums) / float64(len(nums))
}

func med(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	sort.Float64s(nums)
	n := len(nums)
	if n%2 == 0 {
		return (nums[n/2-1] + nums[n/2]) / 2
	}
	return nums[n/2]
}
