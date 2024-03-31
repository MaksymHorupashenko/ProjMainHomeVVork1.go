package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomGenerator(min, max int, ch chan<- int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randomNumber := rand.Intn(max-min+1) + min
		ch <- randomNumber
		time.Sleep(time.Second)
	}
	close(ch)
}

func findMinMax(numbers <-chan int, result chan<- [2]int) {
	min := int(^uint(0) >> 1)
	max := int(^int(0) >> 1)
	for num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	result <- [2]int{min, max}
	close(result)
}

func main() {
	min, max := 1, 100
	numbers := make(chan int)
	result := make(chan [2]int)

	go randomGenerator(min, max, numbers)
	go findMinMax(numbers, result)

	minMax := <-result
	fmt.Printf("Найменше число: %d, Найбільше число: %d\n", minMax[0], minMax[1])
}
