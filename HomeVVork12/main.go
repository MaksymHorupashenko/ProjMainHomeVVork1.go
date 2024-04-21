package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const passwordsFile = "data/passwords.txt"

func main() {
	for {
		fmt.Println("1. Переглянути паролі")
		fmt.Println("2. Зберегти пароль")
		fmt.Println("3. Дістати пароль")
		fmt.Println("4. Вийти")
		var choice int
		fmt.Print("Ваш вибір: ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			viewPasswords()
		case 2:
			savePassword()
		case 3:
			getPassword()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Неправильний вибір. Будь ласка, спробуйте знову.")
		}
	}
}

func viewPasswords() {
	passwords, err := readPasswords()
	if err != nil {
		log.Fatalf("Помилка при читанні файлу паролів: %v", err)
	}

	fmt.Println("Збережені паролі:")
	for service, password := range passwords {
		fmt.Printf("%s: %s\n", service, password)
	}
}

func savePassword() {
	fmt.Print("Введіть назву сервісу або сайту: ")
	serviceName := readLine()
	fmt.Print("Введіть пароль: ")
	password := readLine()

	file, err := os.OpenFile(passwordsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Помилка при відкритті файлу паролів: %v", err)
	}
	defer file.Close()

	if _, err := fmt.Fprintf(file, "%s: %s\n", serviceName, password); err != nil {
		log.Fatalf("Помилка при записі паролю в файл: %v", err)
	}

	fmt.Println("Пароль успішно збережено!")
}

func getPassword() {
	fmt.Print("Введіть назву сервісу або сайту: ")
	serviceName := readLine()

	passwords, err := readPasswords()
	if err != nil {
		log.Fatalf("Помилка при читанні файлу паролів: %v", err)
	}

	password, ok := passwords[serviceName]
	if !ok {
		fmt.Println("Пароль для цього сервісу не знайдено.")
		return
	}

	fmt.Printf("Пароль для сервісу %s: %s\n", serviceName, password)
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readPasswords() (map[string]string, error) {
	file, err := os.Open(passwordsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	passwords := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")
		if len(parts) == 2 {
			passwords[parts[0]] = parts[1]
		}
	}

	return passwords, scanner.Err()
}
