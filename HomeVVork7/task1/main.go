package main

import (
	"fmt"
	"time"
)

func randomGenerator(ch chan<- int) {
	for {
		randomNumber := generateRandomNumber()
		ch <- randomNumber
		time.Sleep(time.Second)
	}
}

func generateRandomNumber() int {
	return time.Now().Nanosecond() % 15
}

func averageCalculator(input <-chan int, output chan<- float64) {
	sum := 0
	count := 0
	for randomNumber := range input {
		sum += randomNumber
		count++
		average := float64(sum) / float64(count)
		output <- average
	}
}

func printer(output <-chan float64, done chan<- bool) {
	for average := range output {
		fmt.Printf("Середнє значення: %.2f\n", average)
	}
	done <- true 
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan float64)
	done := make(chan bool) 

	go randomGenerator(ch1)
	go averageCalculator(ch1, ch2)
	go printer(ch2, done)

	<-done
}
