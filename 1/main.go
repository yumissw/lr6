package main

import (
	"fmt"
	"math/rand"
	"time"
)

func factorial(n int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Printf("факториал %d = %d\n", n, result)
}

func randomNumbers(count int) {
	for i := 0; i < count; i++ {
		num := rand.Intn(100)
		fmt.Printf("%d-ое случайное число: %d\n", i+1, num)
	}
}

func sumSeries(n int) {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("сумма ряда от 1 до %d = %d\n", n, sum)
}

func main() {

	go factorial(4)
	go randomNumbers(5)
	go sumSeries(10)

	time.Sleep(2 * time.Second)
}
