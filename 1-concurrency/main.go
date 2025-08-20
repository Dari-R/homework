package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch1, ch2 := make(chan int, 10), make(chan int, 10)
	var wg1, wg2 sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go func() {
			makeSlice(ch1)
			wg1.Done()
		}()
	}
	go func() {
		wg1.Wait()
		close(ch1)
	}()
	for v := range ch1 {
		wg2.Add(1)
		go func(x int) {
			sqrtFunc(x, ch2)
			wg2.Done()
		}(v)

	}
	go func() {
		wg2.Wait()
		close(ch2)

	}()
	for v := range ch2 {
		fmt.Println(v)
	}

}

func makeSlice(ch chan int) {
	a := rand.Intn(100) + 1
	ch <- a
}

func sqrtFunc(res int, ch2 chan int) {
	ch2 <- (res * res)
}
