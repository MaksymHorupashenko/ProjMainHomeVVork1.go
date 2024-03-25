package main

import (
	"fmt"
	"os"
)

type Player struct {
	Name     string
}

type Situation struct {
	ID          string
	Description string
	Options     []Option
}

type Option struct {
	Text   string
	Action func(*Player) Situation
}

func main() {
	fmt.Println("Ласкаво просимо до гри «Новий світ»!")

	player := Player{Name: "Стівен"}
	currentSituation := startSituation()

	for {
		fmt.Println(currentSituation.Description)
		fmt.Println("Оберіть дію:")

		for i, option := range currentSituation.Options {
			fmt.Printf("%d. %s\n", i+1, option.Text)
		}

		var choice int
		fmt.Print("> ")
		fmt.Scanln(&choice)

		if choice < 1 || choice > len(currentSituation.Options) {
			fmt.Println("Невірний вибір. Спробуйте ще раз.")
			continue
		}

		var action = currentSituation.Options[choice-1].Action
		currentSituation = action(&player)
	}
}

func startSituation() Situation {
	return Situation{
		ID:          "start",
		Description: "Стівен прокинувся біля входу в печеру. Він лише памʼятає своє імʼя.  У печері темно, тому Стівен вирішує не йти до неї.",
		Options: []Option{
			{
				Text:   "Йти в ліс",
				Action: goIntoForest,
			},
			{
				Text:   "Йти до озера",
				Action: goIntoLake,
			},
		},
	}
}

func goIntoForest(player *Player) Situation {
	fmt.Println("Стівен натикається на мертве тіло дивної тварини біля неї лежала табличка на якій було написано 42. Він обирає нічого з ним не робити й іти далі.")
	return Situation{
		ID:          "forest",
		Description: "Через деякий час Стівен приходить до безлюдного табору. У найближчому наметі він знаходить сейф з кодовим замком з двох чисел.",
		Options: []Option{
			{
				Text:   "Спробувати відкрити сейф",
				Action: tryOpenSafe,
			},
			{
				Text:   "Відпочити",
				Action: restAtCamp,
			},
		},
	}
}

func goIntoLake(player *Player) Situation {
	fmt.Println("Стівен прямує дорогою до озера. Він підслизнувся і впав у озеро. Нажаль він не вміє плавати.")
	return startSituation()
}

func restAtCamp(player *Player) Situation {
	fmt.Println("Cтівен вирішує відпочити, він присідає і в цей час його кусає комаха. Стівен непритомніє.")
	fmt.Println("Гра завершена.")
	os.Exit(0)
	return Situation{}
}

func tryOpenSafe(player *Player) Situation {
	fmt.Println("Стівен спробує відкрити сейф.")
	fmt.Println("Введіть код:")
	var code int
	fmt.Scanln(&code)
	if code == 42 {
		fmt.Println("Сейф відкрито! У сейфі знаходиться карта, Стівен вирішує слідувати ній і знаходить рятівників. Вітаю, Гра завершена.")
		os.Exit(0)
	} else {
		fmt.Println("Сейф залишається закритим. Стівен шукає інший вихід.")
		return startSituation()
	}
	return startSituation()
}
