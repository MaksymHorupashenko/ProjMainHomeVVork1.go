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

func (p *Player) play(round Round) {
	fmt.Printf("[%s] Question: %s\n", p.Name, round.Question)
	for i, option := range round.Options {
		fmt.Printf("[%s] Option %d: %s\n", p.Name, i+1, option)
	}

	// Просто для прикладу, гравець вибирає випадкову відповідь
	ans := rand.Intn(len(round.Options)) + 1
	if ans == round.Answer {
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

	// Обробка сигналів OS
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Println("\nReceived termination signal. Exiting...")
		cancel()
	}()

	// Гравці
	players := []*Player{
		newPlayer("Player1"),
		newPlayer("Player2"),
		newPlayer("Player3"),
	}

	// Генерація раундів
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				round := generateRound()
				for _, player := range players {
					player.play(round)
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	<-ctx.Done() // Блокуємо головний потік доки не буде скасований контекст
	fmt.Println("Гра завершена!")
}


