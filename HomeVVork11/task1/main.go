package main


import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	content, err := ioutil.ReadFile("numbers.txt")
	if err != nil {
		fmt.Println("Помилка при зчитуванні файлу:", err)
		return
	}

	phoneRegex := regexp.MustCompile(`(\+?(\d{1,3})[ -]?)?(\(?\d{3}\)?[ -]?)?\d{3}[ -]?\d{4}`)

	phoneNumbers := phoneRegex.FindAllString(string(content), -1)

	fmt.Println("Знайдені номери телефонів:")
	for _, number := range phoneNumbers {
		fmt.Println(number)
	}
}
