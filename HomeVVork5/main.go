package main

import (
	"fmt"
	"strings"
  
type TextEditor struct {
	texts []string
	index map[string][]int
}

func (editor *TextEditor) LoadTexts(texts []string) {
	editor.texts = texts
	editor.index = editor.buildIndex()
}

func (editor *TextEditor) buildIndex() map[string][]int {
	index := make(map[string][]int)
	for i, text := range editor.texts {
		words := strings.Fields(text)
		uniqueWords := make(map[string]bool)
		for _, word := range words {
			if _, found := uniqueWords[word]; !found {
				index[word] = append(index[word], i)
				uniqueWords[word] = true
			}
		}
	}
	return index
}

func (editor *TextEditor) FindTextsByWord(word string) []string {
	if editor.index == nil {
		return nil
	}

	indices, ok := editor.index[word]
	if !ok {
		return nil
	}

	var foundTexts []string
	uniqueTexts := make(map[string]bool)
	for _, idx := range indices {
		text := editor.texts[idx]
		if _, found := uniqueTexts[text]; !found {
			foundTexts = append(foundTexts, text)
			uniqueTexts[text] = true
		}
	}
	return foundTexts
}

func main() {
	editor := TextEditor{}

	texts := []string{
		"Hash map data structures use a hash  function, which turns a key into an index within an underlying array.",
		"The hash function can be used to access an index when inserting a value or retrieving a value from a hash map.",
		"Hash map underlying data structure", "Hash maps are built on top of an underlying array data structure using an indexing system.",
		"Each index in the array can store one key-value pair.",
		"If the hash map is implemented using chaining for collision resolution, each index can store another data structure such as a linked list, which stores all values for multiple keys that hash to the same index.",
		"Each Hash Map key can be paired with only one value. However, different keys can be paired with the same value.",
	}

	editor.LoadTexts(texts)

	foundTexts := editor.FindTextsByWord("hash")

	fmt.Println("Рядки, які містять слово 'hash':")
	for _, text := range foundTexts {
		fmt.Println(text)
	}
}
  
