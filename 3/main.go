package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generate(ch chan<- int, uniqueNumbers *sync.Map) {
	for {
		num := rand.Intn(15)
		if _, loaded := uniqueNumbers.LoadOrStore(num, struct{}{}); !loaded {
			ch <- num // уникальное число отправляем в канал
		}
		//time.Sleep(time.Second)
	}
}

func check(ch <-chan int, resultCh chan<- string) {
	for num := range ch {
		if num%2 == 0 {
			resultCh <- fmt.Sprintf("%d четное", num)
		} else {
			resultCh <- fmt.Sprintf("%d нечетное", num)
		}
	}
}

func main() {
	numberCh := make(chan int)
	resultCh := make(chan string)

	var uniqueNumbers sync.Map // Используем sync.Map для хранения уникальных чисел

	go generate(numberCh, &uniqueNumbers)
	go check(numberCh, resultCh)

	for {
		select {
		case result := <-resultCh: // принимаем сообщения о чётности/нечётности
			fmt.Println(result)
		case <-time.After(2 * time.Second):
			return
		}
	}
}
