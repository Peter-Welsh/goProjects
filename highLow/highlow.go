/*
This program is a simple game of high-low, in which the user tries to guess a number chosen at random.
The user can specify the range of possible numbers by passing the min and max command line arguments.
The program keeps track of your running average number of guesses and displays it to the screen.
There is also an autopilot mode, which works like a binary search algorithm, and you can specify how many times it runs.

Arguments:
-min (The lower end of the range of possible answers)
-max (The upper end of the range of possible answers)
-a (auto-pilot mode; the program enters the guesses automatically)
-i (number of iterations to go through in auto-pilot mode before prompting to continue)
-f (fast mode; removes the delay between guesses in auto-pilot mode)

e.g. to simulate a coin flipper that flips 1000 coins, "go run highlow.go -min=0 -max=1 -a -i=1000 -f"
or to find the average number of iterations a binary search algorithm might take to find a certain
value in a sorted array of 100 elements, "go run highlow.go -min=1 -max=100 -a -i=1000 -f" (it's about 5.8)

By Peter Welsh
09/06/2022
*/

package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	min, max, autopilot, iterations, fastMode, shouldReturn := parseArguments()
	if shouldReturn {
		return
	}
	timesPlayed, totalGuesses := 0, 0 //used to calculate the running average
	i := iterations
	for {
		timesPlayed++
		playHighLow(max, min, autopilot, &timesPlayed, &totalGuesses, fastMode)
		i--
		if i == 0 {
			time.Sleep(10 * time.Millisecond)
			var continueOrQuit string
			fmt.Println("Press enter to restart or Q to quit.")
			fmt.Scanln(&continueOrQuit)
			if strings.EqualFold(continueOrQuit, "Q") {
				break
			}
			i = iterations //reset number of iterations
		}
	}
}

func parseArguments() (int, int, bool, int, bool, bool) {
	var min int
	var max int
	var autopilot bool
	var iterations int
	var fastMode bool
	flag.IntVar(&min, "min", 1, "Specify the minimum end of the range (inclusive) of possible numbers. Default is 1.")
	flag.IntVar(&max, "max", 100, "Specify the maximum end of the range (inclusive) of possible numbers. Default is 100.")
	flag.BoolVar(&autopilot, "a", false, "Specify whether to turn on auto-pilot mode. Default is false.")
	flag.IntVar(&iterations, "i", 1, "Specify how many iterations to run through in auto-pilot mode. Default is 1.")
	flag.BoolVar(&fastMode, "f", false, "Specify whether to make auto-pilot mode go fast. Default is false.")
	flag.Parse()
	if min > max {
		fmt.Println("Max must not be less than min. Exiting.")
		return 0, 0, false, 0, false, true
	}
	if iterations <= 0 {
		fmt.Println("Iterations must be a positive integer. Exiting.")
		return 0, 0, false, 0, false, true
	}
	if !autopilot {
		fastMode = false
		iterations = 1
	}
	return min, max, autopilot, iterations, fastMode, false
}

func playHighLow(max int, min int, autopilot bool, timesPlayed *int, totalGuesses *int, fastMode bool) {
	possibilities := max - min + 1
	lowerEndOfRange := min
	upperEndOfRange := max
	// par is the number of guesses in the worst case scenario when playing intelligently (guessing midpoints)
	par := int(math.Ceil(math.Log2(float64(possibilities))))
	// Intn(n) gives a number in the range [0,n)
	// i.e. Intn(99) will give numbers 0 through 98
	// Intn(100) + 1 will give us the desired range of 1 thru 100 inclusive when min=1 and max=100
	answer := rand.Intn(possibilities) + min
	guesses := 0
	fmt.Printf("I'm thinking of a number between %d and %d.\r\n", min, max)
	for {
		guess := getNumberInput(autopilot, lowerEndOfRange, upperEndOfRange, fastMode)
		guesses++
		if guess > answer {
			fmt.Println("No, lower.")
			upperEndOfRange = guess - 1
		} else if guess < answer {
			fmt.Println("No, higher.")
			lowerEndOfRange = guess + 1
		} else {
			*totalGuesses += guesses
			averageGuesses := float32(*totalGuesses) / float32(*timesPlayed)
			successfulGuess(guesses, par, averageGuesses)
			break
		}
	}
}

func successfulGuess(guesses int, par int, averageGuesses float32) {
	only := ""
	es := ""
	if guesses > par {
		fmt.Print("You finally got it. ")
	} else if guesses < par {
		fmt.Print("Nicely done! ")
		only = "only "
	} else {
		fmt.Print("Not bad. ")
	}
	if guesses != 1 {
		es = "es"
	}
	fmt.Printf("It took you %s%d guess%s.\r\n", only, guesses, es)
	fmt.Printf("Your running average is %f.\r\n\r\n", averageGuesses)
}

func getNumberInput(autopilot bool, lowerEndOfRange int, upperEndOfRange int, fastMode bool) int {
	var number int
	if autopilot {
		number = (upperEndOfRange + lowerEndOfRange) / 2
		wait(fastMode)
		fmt.Printf("%d?\r\n", number)
		wait(fastMode)
	} else {
		fmt.Print("What's your guess? ")
		fmt.Scanln(&number)
	}
	return number
}

func wait(fastMode bool) {
	if !fastMode {
		time.Sleep(500 * time.Millisecond)
	}
}
