package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, ch chan<- int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a //запись в канал
		a, b = b, a+b
	}
	close(ch)
}

func print(ch <-chan int) {
	for num := range ch { // чтение из канала
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int)

	go fibonacci(10, ch)
	go print(ch)

	time.Sleep(time.Second)
}
