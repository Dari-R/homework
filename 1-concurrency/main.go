package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch1, ch2 := make(chan []int), make(chan int, 10)
	go func() {
		makeSlice(ch1)
		close(ch1)
	}()

	go func() {
		sl:=<-ch1
		for _, s := range sl {
			sqrtFunc(s, ch2)
		}

		close(ch2)
	}()

	for v := range ch2 {
		fmt.Println(v)
	}

}

func makeSlice(ch chan []int) {
	sl := make([]int, 10)
	for i := 0; i < 10; i++ {
		a := rand.Intn(100) + 1
		sl[i] = a
	}
	ch <- sl
}

func sqrtFunc(res int, ch2 chan int) {
	ch2 <- res * res
}
