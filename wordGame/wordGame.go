/*
This program is a word game similar to Wordle.
The user enters the length of the word, and the program generates a random word of that length.
The user then guesses words until they get it correct.
After each guess, the program displays clues to help the user arrive at the answer.
A letter in the correct spot will yield a green square
A letter in the word but not in the right spot will yield a yellow square.
A letter that is not in the word at all will yield a black square.

Note:
- The program needs to be run in a console that prints in a font that supports full-color UTF8 characters
- The program was developed and tested on the VS Code terminal
- To give up, you can type !q when prompted for a guess
- To play a game with a word length other than 5, pass the length as a command line argument like so: "go run wordgame.go -l x" (where x is a positive integer)

Limitations:
- The program relies on a specific external API to get the hidden word and will not function if that API is unavailable.
- The program relies on a dictionary API (defined in ../dictionary) which may not have the definition

Possible improvements:
- add a graphical keyboard to show which letters have been eliminated
- keep track of best scores
- check if the user's guess is a real word rather than allowing any random letters
- make the word length a command-line argument
- support consoles that don't have unicode support by displaying clues in ASCII (e.g. G/Y/B for Green Yellow Black)

By Peter Welsh
08/31/2022
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/exp/slices"

	dict "example.com/dictionary"
)

const baseUrl = "https://random-word-api.herokuapp.com/word?length="

func main() {
	var wordLength int
	flag.IntVar(&wordLength, "l", 5, "Specify word length. Default is 5.")
	flag.Parse()
	if wordLength < 2 || wordLength > 15 {
		fmt.Println("Invalid parameter value.\r\nThe word length must be between 2 and 15.\r\nExiting.")
		return
	}
	for {
		playWordGame(wordLength)

		var continueOrQuit string
		fmt.Println("\r\nPress enter to play again or Q to quit.")
		fmt.Scanln(&continueOrQuit)
		if strings.EqualFold(continueOrQuit, "Q") {
			break
		}
	}
}

var quitCommands = []string{"!q", "!quit", "!exit"}

const commandRune = '!'

func playWordGame(wordLength int) {
	answer := getAnswer(wordLength)
	guesses := 0
	for {
		guess := getGuess(wordLength)
		if guess[0] == commandRune {
			if slices.Contains(quitCommands, guess) {
				fmt.Printf("The word was %q.\r\n", answer)
				break
			} else {
				fmt.Println("Unrecognized command.")
				continue
			}
		}
		guesses++
		correct := checkGuess(guess, answer)
		if correct {
			pluralEnding := "es"
			if guesses == 1 {
				pluralEnding = ""
			}
			fmt.Printf("Well done! You got it in %d guess%s.\r\n", guesses, pluralEnding)
			break
		}
	}
	time.Sleep(time.Second)
	fmt.Printf("Would you like to see the definition of %q? (Y/N) ", answer)
	var yesOrNo string
	fmt.Scanln(&yesOrNo)
	if strings.EqualFold(yesOrNo, "Y") {
		dict.DefineWord(answer)
	}
}

func getAnswer(length int) string {
	url := fmt.Sprintf(baseUrl+"%d", length)
	resp, err1 := http.Get(url)
	if err1 != nil {
		panic(err1)
	}
	answer := readResponse(resp)
	if len(answer) != length {
		panic("Fatal error. Failed to parse the response from the service.")
	}
	return answer
}

func checkGuess(guess string, answer string) bool {
	clue := getClue(answer, guess)
	fmt.Printf("%s\r\n", clue)
	return strings.EqualFold(guess, answer)
}

func getClue(answer string, guess string) string {
	clue := strings.Repeat("⬛", len(answer))
	alreadyMarked := make([]bool, len(answer))
	tmpAnswer := answer
	const placeholder = '~'
	// 1st pass: mark the correct letters and replace the letter with the placeholder
	for i, letter := range guess {
		if letter == rune(answer[i]) {
			replaceRuneAtIndex(&clue, i, '🟩')
			replaceRuneAtIndex(&tmpAnswer, i, placeholder)
			alreadyMarked[i] = true
		}
	}
	// 2nd pass: mark the remaining letters that are part of the answer
	for i, letter := range guess {
		if alreadyMarked[i] {
			// do not overwrite a green square with a yellow
			continue
		}
		if strings.Contains(tmpAnswer, string(letter)) {
			replaceRuneAtIndex(&clue, i, '🟨')
			replaceRuneAtIndex(&tmpAnswer, strings.IndexRune(tmpAnswer, letter), placeholder)
		}
	}
	return clue
}

func replaceRuneAtIndex(word *string, i int, letter rune) {
	tmpWord := []rune(*word)
	tmpWord[i] = letter
	*word = string(tmpWord)
}

func readResponse(resp *http.Response) string {
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("The server isn't responding. Try again later.")
		return ""
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// the response body will look like this: ["word"]
	const start = "[\""
	const end = "\"]"
	word := bodyBytes[len(start) : len(bodyBytes)-len(end)]
	return string(word)
}

func getGuess(length int) string {
	var guess string
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)
		if len(strings.TrimSpace(guess)) == length || (len(guess) > 0 && guess[0] == commandRune) {
			break
		}
		fmt.Printf("Your guess should be %d characters long. Try again.\r\n", length)
	}
	return guess
}
