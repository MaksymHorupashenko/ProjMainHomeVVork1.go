package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
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
	Score int
}

func newPlayer(name string) *Player {
	return &Player{
		Name: name,
	}
}

func (p *Player) play(round Round, answer int) {
	if answer == round.Answer {
		fmt.Printf("[%s] Correct answer!\n", p.Name)
		p.Score++
	} else {
		fmt.Printf("[%s] Wrong answer!\n", p.Name)
	}
}

func generateRound() Round {
	questions := []string{"Який футбольний клуб з Англії?", "Хто виграв Золотий мяч?", "У якому клубі грав Андрій Шевченко?"}
	options := [][]string{
		{"Челсі", "Реал Мадрид", "Шахтар", "Ювентус"},
		{"Шевченко", "Мілевський", "Мудрик", "Зінченко"},
		{"Арсенал", "Динамо Київ", "Шахтар", "Барселона"},
	}
	answers := []int{1, 1, 2}

	idx := rand.Intn(len(questions))
	return Round{
		Question: questions[idx],
		Options:  options[idx],
		Answer:   answers[idx],
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Println("\nОтриманий сигнал завершення.")
		cancel()
	}()

	players := []*Player{
		newPlayer("Player1"),
		newPlayer("Player2"),
		newPlayer("Player3"),
	}

	answerCh := make(chan int)
	defer close(answerCh)

	resultCh := make(chan string)
	defer close(resultCh)

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				round := generateRound()
				fmt.Printf("\nНовий раунд:\nПитання: %s\n", round.Question)
				for i, option := range round.Options {
					fmt.Printf("Option %d: %s\n", i+1, option)
				}

				for _, player := range players {
					answer := rand.Intn(len(round.Options)) + 1
					player.play(round, answer)
				}

				time.Sleep(3 * time.Second)
				
				result := ""
				for _, player := range players {
					result += fmt.Sprintf("[%s: %d] ", player.Name, player.Score)
				}
				resultCh <- result
			case <-ctx.Done():
				fmt.Println("Гра завершена!")
				return
			}
		}
	}()

	for {
		select {
		case result := <-resultCh:
			fmt.Println("Результати раунду:", result)
		case <-ctx.Done():
			fmt.Println("Гра завершена!")
			return
		}
	}
}

