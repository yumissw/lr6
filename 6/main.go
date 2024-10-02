package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Line string
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func worker(tasks <-chan Task) {
	for task := range tasks { //получаем
		reversed := reverse(task.Line)
		fmt.Println(reversed)
	}
}

func main() {
	var numWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&numWorkers)

	tasks := make(chan Task)

	for i := 0; i < numWorkers; i++ { //создание воркеров
		go worker(tasks)
	}

	file, err := os.Open("input.txt") //чтение строк из файла
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks <- Task{Line: scanner.Text()}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения файла:", err)
	}

	file.Close()
	close(tasks) // Закрываем канал задач после завершения чтения
	time.Sleep(50 * time.Millisecond)
}
