package main


import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	content, err := ioutil.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Помилка при зчитуванні файлу:", err)
		return
	}

	wordRegex := regexp.MustCompile(`\b[a-zA-Z]*[aeiouAEIOU][a-zA-Z]*[^aeiouAEIOU\W]\b`)

	words := wordRegex.FindAllString(string(content), -1)

	fmt.Println("Знайдені слова, що починаються на голосну та закінчуються на приголосну:")
	for _, word := range words {
		fmt.Println(word)
	}
}

