package main

import (
	"fmt"
	"time"
)

// Структура для запроса
type Request struct {
	Operation string
	A         float64
	B         float64
}

type Response struct {
	Result float64
	Error  error
}

var requestChan = make(chan Request)
var responseChan = make(chan Response)

func calculator() {
	for req := range requestChan {
		var result float64
		var err error

		if req.Operation == "+" {
			result = req.A + req.B
		} else if req.Operation == "-" {
			result = req.A - req.B
		} else if req.Operation == "*" {
			result = req.A * req.B
		} else if req.Operation == "/" {
			if req.B != 0 {
				result = req.A / req.B
			} else {
				err = fmt.Errorf("деление на ноль")
			}
		} else {
			err = fmt.Errorf("неизвестная операция: %s", req.Operation)
		}

		responseChan <- Response{Result: result, Error: err}
	}
}

func main() {
	go func() {
		calculator()
	}()

	requests := []Request{
		{Operation: "+", A: 21, B: 5},
		{Operation: "-", A: 13, B: 8},
		{Operation: "*", A: 37, B: 9},
		{Operation: "/", A: 57, B: 14},
	}

	for _, req := range requests {
		requestChan <- req

		response := <-responseChan
		if response.Error != nil {
			fmt.Printf("Ошибка: %s\n", response.Error)
		} else {
			fmt.Printf("%.2f %s %.2f = %.2f\n", req.A, req.Operation, req.B, response.Result)
		}
	}

	close(requestChan)
	time.Sleep(time.Second)
	close(responseChan)
}
