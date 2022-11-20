/*
This program prompts the user for a word and displays its definition.

Limitations:
- The program depends on a specific external API and will not function if that API is unavailable.
- The input must be 1 word; the program cannot define phrases such as "quantum mechanics"

By Peter Welsh
08-23-2022
*/

package main

import (
	"fmt"
	"strings"
	"time"

	dict "example.com/dictionary"
)

func main() {
	for {
		word := getWordInput()
		dict.DefineWord(word)

		time.Sleep(time.Second)
		var continueOrQuit string
		fmt.Println("\r\nPress enter to restart or Q to quit.")
		fmt.Scanln(&continueOrQuit)
		if strings.EqualFold(continueOrQuit, "Q") {
			break
		}
	}
}

func getWordInput() string {
	var word string
	for {
		fmt.Print("Enter a word to see its definition(s): ")
		fmt.Scanln(&word)
		if word != "" {
			break
		}
		fmt.Print("Try again. ")
	}
	return word
}
