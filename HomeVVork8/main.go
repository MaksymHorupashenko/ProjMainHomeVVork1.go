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

func (p *Player) play(round Round, answerCh chan int) {
	fmt.Printf("[%s] Question: %s\n", p.Name, round.Question)
	for i, option := range round.Options {
		fmt.Printf("[%s] Option %d: %s\n", p.Name, i+1, option)
	}

	var ans int
	select {
	case ans = <-answerCh:
		if ans == round.Answer {
			fmt.Printf("[%s] Correct answer!\n", p.Name)
			p.Score++
		} else {
			fmt.Printf("[%s] Wrong answer!\n", p.Name)
		}
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
		fmt.Println("\nReceived termination signal. Exiting...")
		cancel()
	}()

	players := []*Player{
		newPlayer("Player1"),
		newPlayer("Player2"),
		newPlayer("Player3"),
	}

	answerChs := make([]chan int, len(players))
	for i := range answerChs {
		answerChs[i] = make(chan int)
	}

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				round := generateRound()
				for i, player := range players {
					go func(p *Player, round Round, ch chan int) {
						p.play(round, ch)
					}(player, round, answerChs[i])
				}
				time.Sleep(10 * time.Second) 
				for _, ch := range answerChs {
					close(ch) 
				}
				answerChs = make([]chan int, len(players)) 
				for i := range answerChs {
					answerChs[i] = make(chan int)
				}
			case <-ctx.Done():
				fmt.Println("Гра завершена!")
				return
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Гра завершена!")
			return
		default:
			var input int
			fmt.Println("Введіть ваш варіант відповіді:")
			fmt.Scanln(&input)
			for _, ch := range answerChs {
				ch <- input
			}
		}
	}
}


