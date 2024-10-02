package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	counter := 0
	const num = 3
	const inc = 10

	for i := 0; i < num; i++ {
		go func() {
			for j := 0; j < inc; j++ {
				mu.Lock() // блокировка мьютекса перед изменением счетчика
				counter++
				fmt.Println("счетчик - ", counter)
				mu.Unlock() // освобождение мьютекса
			}
		}()
	}

	time.Sleep(10 * time.Second)

	fmt.Printf("результат: %d\n", counter)
}
