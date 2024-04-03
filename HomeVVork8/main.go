package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Round struct {
	Question string
	Options  []string
	Answer   int
}

type Player struct {
	Name  string
	Input chan int
	Score int
}

func newPlayer(name string) *Player {
	return &Player{
		Name:  name,
		Input: make(chan int),
	}
}

func (p *Player) play(ctx context.Context, round Round, resultChan chan<- string) {
	fmt.Printf("[%s] Question: %s\n", p.Name, round.Question)
	for i, option := range round.Options {
		fmt.Printf("[%s] Option %d: %s\n", p.Name, i+1, option)
	}

	select {
	case ans := <-p.Input:
		if ans == round.Answer {
			fmt.Printf("[%s] Correct answer!\n", p.Name)
			p.Score++
		} else {
			fmt.Printf("[%s] Wrong answer!\n", p.Name)
		}
	case <-ctx.Done():
		fmt.Printf("[%s] Game ended!\n", p.Name)
		resultChan <- fmt.Sprintf("[%s] Game ended!\n", p.Name)
	}
}

func generateRound() (Round, Round, Round) {
	questions := []string{"Який футбольний клуб з Англії?", "Хто виграв Золотий мяч?", "У якому клубі грав Андрій Шевченко?"}
	options := [][]string{
		{"Челсі", "Реал Мадрид", "Шахтар", "Ювентус"},
		{"Шевченко", "Мілевський", "Мудрик", "Зінченко"},
		{"Арсенал", "Динамо Київ", "Шахтар", "Барселона"},
	}
	answers := []int{1, 1, 2}

	idx1 := rand.Intn(len(questions))
	idx2 := (idx1 + 1) % len(questions)
	idx3 := (idx1 + 2) % len(questions)

	round1 := Round{
		Question: questions[idx1],
		Options:  options[idx1],
		Answer:   answers[idx1],
	}
	round2 := Round{
		Question: questions[idx2],
		Options:  options[idx2],
		Answer:   answers[idx2],
	}
	round3 := Round{
		Question: questions[idx3],
		Options:  options[idx3],
		Answer:   answers[idx3],
	}

	return round1, round2, round3
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Println("\nReceived termination signal. Exiting...")
		cancel()
	}()

	var wg sync.WaitGroup

	rounds := make(chan Round)
	resultChan := make(chan string)
 
	players := []*Player{
		newPlayer("Player1"),
		newPlayer("Player2"),
		newPlayer("Player3"),
	}

	for _, player := range players {
		wg.Add(1)
		go func(p *Player) {
			defer wg.Done()
			for {
				select {
				case round := <-rounds:
					p.play(ctx, round, resultChan)
				case <-ctx.Done():
					return
				}
			}
		}(player)
	}


	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				round1, round2, round3 := generateRound()
				for _ = range players {
					select {
					case rounds <- round1:
					case <-ctx.Done():
						return
					}
					select {
					case rounds <- round2:
					case <-ctx.Done():
						return
					}
					select {
					case rounds <- round3:
					case <-ctx.Done():
						return
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case result := <-resultChan:
				fmt.Println(result)
			}
		}
	}()

	wg.Wait()
	fmt.Println("Гра завершена!")
}
