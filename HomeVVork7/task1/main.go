package main

import (
	"fmt"
	"sync"
	"time"
)

func randomGenerator(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		randomNumber := generateRandomNumber()
		ch <- randomNumber
		time.Sleep(time.Second)
	}
}

func generateRandomNumber() int {
	return time.Now().Nanosecond() % 15
}

func averageCalculator(input <-chan int, output chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	count := 0
	for randomNumber := range input {
		sum += randomNumber
		count++
		average := float64(sum) / float64(count)
		output <- average
	}
}

func printer(output <-chan float64) {
	for average := range output {
		fmt.Printf("Середнє значення: %.2f\n", average)
	}
}

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan float64)

	wg.Add(1)
	go randomGenerator(ch1, &wg)

	wg.Add(1)
	go averageCalculator(ch1, ch2, &wg)

	go printer(ch2)

	wg.Wait()
}
