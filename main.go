// Author: Bianca Boo
// Version: 1.0.0
// Date: 2025-12-10
// Fileoverview: This program uses a word search in a given sentence.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// function returns true if word found
func findWord(sentence string, word string) bool {
	return strings.Contains(sentence, word)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	// get sentence
	fmt.Println("Please enter a sentence?")
	sentence, _ := reader.ReadString('\n')
	sentence = strings.TrimSpace(sentence)

	// get word
	fmt.Println("Please enter a word to search for in your sentence?")
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)

	// result
	if findWord(sentence, word) {
		fmt.Println("Hooray, the word", word, "was found in the sentence:", sentence)
	} else {
		fmt.Println("Sorry, the word", word, "was not found in the sentence:", sentence)
	}

}