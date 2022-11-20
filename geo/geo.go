/*
This program quizzes the user on the capitals of countries around the world.

The user can exit at any time by entering "quit".
The user can also get a hint by entering "hint".
The answer checker has a tolerance for typos; answers like "Port Villa" will be accepted for the capital "Port Vila"
At the end, the user is graded on how well they did.

The program makes a call to an external API if it cannot find a specific json file from which to read the countries and their capitals.
After making the API call the first time, it saves the response in a file to avoid making unneeded calls.
If the file is older than a set amount of time (30 days), it makes the call again to refresh the list.

Arguments:
-q (int; Quiz size / Number of questions; default is max)
	-- Example: go run geo.go -q=5
-m (int; Quiz mode; governs the type of countries that appear on the quiz; 0 = Independent, 1 = Dependent, 2 = All; default is 0)
	-- Example: go run geo.go -m=1
-r (bool; Reverse quiz; the user guesses the country given the capital)
    -- Example: go run geo.go -r

Note:
According to restcountries.com, only South Africa has multiple capitals.
Wikipedia disagrees. https://en.wikipedia.org/wiki/List_of_countries_with_multiple_capitals

Limitations:
- It is strict on some answers. E.g., it will mark USA as incorrect when the answer is United States

By Peter Welsh
09/07/2022
My best results below:
You got 193 out of 194 capitals correct (99%) in 9 minutes and 14 seconds (3 seconds per question) and used 0 hints.
Final grade: A+

You got 192 out of 194 countries correct (99%) in 18 minutes and 18 seconds (6 seconds per question) and used 0 hints.
Final grade: A+
*/

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Countries []struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Independent bool     `json:"independent,omitempty"`
	Capitals    []string `json:"capital"`
}

const url = "https://restcountries.com/v3.1/all?fields=name,capital,independent"
const fileName = "countries.json"
const editDistanceTolerance = 1
const doubleEnterPressToleranceMs = 300
const daysTillListExpires = 30

var transformer transform.Transformer

type CountryMode int8

const (
	Independent CountryMode = iota
	Dependent
	All
)

func main() {
	quizSize, countryMode, reverse, shouldReturn := parseArguments()
	if shouldReturn {
		return
	}
	fmt.Println("Initializing")
	// initialize a transformer for the purpose of normalizing letters like the é in Malé
	transformer = transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	start := time.Now()
	numCorrect, numAttempted, numHints := runTheQuiz(quizSize, countryMode, reverse)
	elapsed := time.Since(start)

	printGrade(numCorrect, numAttempted, elapsed, numHints, reverse)
}

func parseArguments() (int, CountryMode, bool, bool) {
	var quizSize int
	var mode int
	var reverse bool
	flag.IntVar(&quizSize, "q", math.MaxInt, "Specify the number of questions. Default is all.")
	flag.IntVar(&mode, "m", 0, "Specify the type of countries to be quizzed on. 0 = Independent; 1 = Depedent; 2 = All. Default is 0.")
	flag.BoolVar(&reverse, "r", false, "Specify whether to reverse the quiz (guess the country given the capital).")
	flag.Parse()
	if quizSize < 1 {
		fmt.Println("Quiz size must be a positive integer. Exiting.")
		return 0, 0, false, true
	}
	if mode < 0 || mode > 2 {
		fmt.Println("Country mode must be a number from 0 to 2. Exiting.")
		return 0, 0, false, true
	}
	countryMode := CountryMode(mode)
	return quizSize, countryMode, reverse, false
}

func runTheQuiz(quizSize int, countryMode CountryMode, reverse bool) (int, int, int) {
	countries := getAllCountries()
	countries = filterCountries(countries, countryMode)
	quizSize = int(math.Min(float64(len(countries)), float64(quizSize)))

	// seed the generator to get new pseudo-random numbers each time the program is run
	rand.Seed(time.Now().UnixNano())
	fmt.Println("\r\nLet's begin.")
	return takeTheQuiz(quizSize, countries, reverse)
}

func takeTheQuiz(quizSize int, countries Countries, reverse bool) (int, int, int) {
	numCorrect := 0
	hintsUsed := 0
	var numAttempted int
	randomIndices := rand.Perm(len(countries))
	for questionNum, i := range randomIndices[0:quizSize] {
		printQuestion(reverse, questionNum, countries, i)
		guess := getGuess()
		hintNum := 0
		for strings.EqualFold(guess, "hint") {
			printHint(countries, i, hintNum, reverse)
			hintNum++
			guess = getGuess()
		}
		hintsUsed += hintNum
		if strings.EqualFold(guess, "quit") || strings.EqualFold(guess, "exit") {
			break
		}
		numAttempted++
		checkAnswer(countries, i, guess, &numCorrect, reverse)
	}
	return numCorrect, numAttempted, hintsUsed
}

func printQuestion(reverse bool, questionNum int, countries Countries, i int) {
	if reverse {
		fmt.Printf("%d. Which country's capital is %s?\r\n", questionNum+1, countries[i].Capitals[0])
	} else {
		fmt.Printf("%d. What is the capital of %s?\r\n", questionNum+1, countries[i].Name.Common)
	}
}

func getGuess() string {
	var guess string
	start := time.Now()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			guess = scanner.Text()
			elapsed := time.Since(start)
			// don't count an accidental double-enter press as a wrong answer
			if elapsed.Milliseconds() > doubleEnterPressToleranceMs {
				break
			}
		}
	}
	return guess
}

func printHint(countries Countries, i int, hintNum int, reverse bool) {
	if reverse {
		fmt.Printf("It starts with the letter %c.\r\n", countries[i].Name.Common[0])
		return
	}
	numCapitals := len(countries[i].Capitals)
	if numCapitals == 1 {
		fmt.Print("It ")
	} else {
		fmt.Print("One of its capitals ")
	}
	// give the first letter of a different capital (if there are more than 1) for consecutive hints
	x := int(math.Min(float64(hintNum%numCapitals), float64(numCapitals-1)))
	fmt.Printf("starts with the letter %c.\r\n", countries[i].Capitals[x][0])
}

func checkAnswer(countries Countries, i int, guess string, numCorrect *int, reverse bool) {
	correct := isCorrect(guess, countries, i, reverse)
	if correct {
		printCorrect(reverse, numCorrect, guess, countries, i)
		return
	}
	printIncorrect(guess, reverse, countries, i)
}

func printCorrect(reverse bool, numCorrect *int, guess string, countries Countries, i int) {
	numCapitals := len(countries[i].Capitals)
	*numCorrect++
	if reverse {
		fmt.Print("Correct!")
		country := countries[i].Name.Common
		if !strings.EqualFold(guess, country) {
			fmt.Printf(" It is %s.", country)
		}
		fmt.Println("")
		return
	}
	if numCapitals == 1 {
		fmt.Print("Correct!")
		capital := countries[i].Capitals[0]
		if !strings.EqualFold(guess, capital) {
			fmt.Printf(" It is %s.", capital)
		}
		fmt.Println("")
		return
	}
	if numCapitals > 1 {
		fmt.Print("Yes, that's one of them. ")
		printMultipleCapitals(numCapitals, countries, i)
	}
}

func printIncorrect(guess string, reverse bool, countries Countries, i int) {
	numCapitals := len(countries[i].Capitals)
	if len(guess) > 0 && !strings.EqualFold(guess, "idk") {
		fmt.Print("No. ")
	}
	if reverse {
		fmt.Printf("It's %s.\r\n", countries[i].Name.Common)
		return
	}
	if numCapitals > 1 {
		printMultipleCapitals(numCapitals, countries, i)
	} else {
		fmt.Printf("It's %s.\r\n", countries[i].Capitals[0])
	}
}

func isCorrect(guess string, countries Countries, i int, reverse bool) bool {
	guessNormalized := normalize(guess)
	if reverse {
		country := countries[i].Name.Common
		countryNormalized := normalize(country)
		areEqual := strings.EqualFold(guessNormalized, countryNormalized)
		return areEqual || editDistance(guessNormalized, countryNormalized) <= editDistanceTolerance
	}
	for _, capital := range countries[i].Capitals {
		capitalNormalized := normalize(capital)
		areEqual := strings.EqualFold(guessNormalized, capitalNormalized)
		if areEqual || editDistance(guessNormalized, capitalNormalized) <= editDistanceTolerance {
			return true
		}
	}
	return false
}

func normalize(capital string) string {
	// change accented letters to non-accented (Malé becomes Male)
	normalized, _, _ := transform.String(transformer, capital)
	// accept "San Marino" for capital "City of San Marino"
	normalized = strings.TrimPrefix(normalized, "City of ")
	// accept "Vatican" for capital "Vatican City"
	normalized = strings.TrimSuffix(normalized, " City")
	// accept "St. John's" for capital "Saint John's"
	if strings.HasPrefix(capital, "Saint ") {
		normalized = strings.TrimPrefix(normalized, "Saint ")
		normalized = "St. " + normalized
	}
	// accept "Washington DC" for "Washington, D.C."
	normalized = strings.ReplaceAll(normalized, ",", "")
	normalized = strings.ReplaceAll(normalized, ".", "")
	// accept "Sanaa" for "Sana'a"
	normalized = strings.ReplaceAll(normalized, "'", "")
	// accept "Porto Novo" for "Porto-Novo"
	normalized = strings.ReplaceAll(normalized, "-", " ")
	// accept "Bosnia & Herzegovina" for "Bosnia and Herzegovina"
	normalized = strings.ReplaceAll(normalized, "&", "and")
	// accept "Democratic Republic of the Congo" for "DR Congo"
	normalized = strings.ReplaceAll(normalized, "Democratic Republic", "DR")
	normalized = strings.ReplaceAll(normalized, " of the ", " ")
	// accept "UK" for "United Kingdom"
	normalized = strings.ReplaceAll(normalized, "United Kingdom", "UK")
	// accept "UAE" for "United Arab Emirates"
	normalized = strings.ReplaceAll(normalized, "United Arab Emirates", "UAE")
	// accept "US" for "United States"
	normalized = strings.ReplaceAll(normalized, "United States", "US")
	return normalized
}

func printMultipleCapitals(numCapitals int, countries Countries, i int) {
	fmt.Printf("It has %d capitals: ", numCapitals)
	fmt.Printf("%s\r\n", strings.Join(countries[i].Capitals, ", "))
}

func getLetterGrade(percentage int) string {
	var letterGrade string
	switch percentage / 10 {
	case 10:
		letterGrade = "A+"
	case 9:
		letterGrade = "A"
		appendPlusOrMinus(percentage, &letterGrade)
	case 8:
		letterGrade = "B"
		appendPlusOrMinus(percentage, &letterGrade)
	case 7:
		letterGrade = "C"
		appendPlusOrMinus(percentage, &letterGrade)
	case 6:
		letterGrade = "D"
		appendPlusOrMinus(percentage, &letterGrade)
	default:
		letterGrade = "F"
	}
	return letterGrade
}

func appendPlusOrMinus(percentage int, letterGrade *string) {
	if percentage%10 >= 7 {
		*letterGrade += "+"
		return
	}
	if percentage%10 <= 2 {
		*letterGrade += "-"
		return
	}
}

func getAllCountries() Countries {
	file, err := os.Stat(fileName)
	var countries Countries
	if err != nil {
		fmt.Println("File " + fileName + " was not found.")
		makeRequest(&countries)
		return countries
	}

	if isExpired(file.ModTime()) {
		fmt.Println("File " + fileName + " was found but is old. Refreshing.")
		makeRequest(&countries)
		return countries
	}

	fmt.Println("Reading file " + fileName)
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(fileContents, &countries)
	if err != nil {
		panic(err)
	}

	return countries
}

func isExpired(t time.Time) bool {
	days := 24 * time.Hour
	return time.Since(t) > daysTillListExpires*days
}

func makeRequest(countries *Countries) {
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Making a " + method + " request to " + url)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &countries)
	if err != nil {
		panic(err)
	}
	writeToFile(body)
}

func writeToFile(body []byte) {
	file, err := os.Create("./" + fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println("Writing the response to " + fileName)
	file.WriteString(string(body))
}

func filterCountries(countries Countries, countryMode CountryMode) Countries {
	var filteredCountries Countries
	for _, country := range countries {
		shouldAppend := false
		// discard "countries" like Antarctica
		hasCapitals := len(country.Capitals) > 0
		switch countryMode {
		case Independent:
			shouldAppend = country.Independent && hasCapitals
		case Dependent:
			shouldAppend = !country.Independent && hasCapitals
		case All:
			shouldAppend = hasCapitals
		}
		if shouldAppend {
			filteredCountries = append(filteredCountries, country)
		}
	}
	return filteredCountries
}

func printGrade(numCorrect int, numAttempts int, elapsed time.Duration, numHints int, reverse bool) {
	percentage := getPercentage(numAttempts, numCorrect)
	subject := getQuizSubject(reverse)
	fmt.Printf("You got %d out of %d %s correct (%d%%) ", numCorrect, numAttempts, subject, percentage)
	printElapsedTime(elapsed, numAttempts)
	hS := getPluralEnding(numHints)
	fmt.Printf("and used %d hint%s.\r\n", numHints, hS)
	letterGrade := getLetterGrade(percentage)
	fmt.Printf("Final grade: %s ", letterGrade)
	if percentage == 100 {
		fmt.Print("⭐")
	}
}

func getQuizSubject(reverse bool) string {
	if reverse {
		return "countries"
	}
	return "capitals"
}

func printElapsedTime(elapsed time.Duration, numAttempts int) {
	minutes := int(elapsed.Minutes())
	seconds := int(math.Round(elapsed.Seconds())) % 60
	perQuestion := getSecondsPerQuestion(numAttempts, elapsed)
	mS := getPluralEnding(minutes)
	sS := getPluralEnding(seconds)
	pqS := getPluralEnding(perQuestion)
	fmt.Printf("in %d minute%s and %d second%s (%d second%s per question) ", minutes, mS, seconds, sS, perQuestion, pqS)
}

func getSecondsPerQuestion(numAttempts int, elapsed time.Duration) int {
	var perQuestion int
	if numAttempts == 0 {
		perQuestion = 0
	} else {
		perQuestion = int(math.Round(elapsed.Seconds() / float64(numAttempts)))
	}
	return perQuestion
}

func getPercentage(numAttempts int, numCorrect int) int {
	var percentage int
	if numAttempts == 0 {
		percentage = 0
	} else {
		percentage = int(math.Round(100 * float64(numCorrect) / float64(numAttempts)))
	}
	return percentage
}

func getPluralEnding(unit int) string {
	if unit == 1 {
		return ""
	}
	return "s"
}

// implementation of the Levenshtein edit distance algorithm from:
// https://www.golangprograms.com/golang-program-for-implementation-of-levenshtein-distance.html
// (modified to be case-insensitive)
func editDistance(str1, str2 string) int {
	s1len := len(str1)
	s2len := len(str2)
	column := make([]int, len(str1)+1)

	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= s1len; y++ {
			oldkey := column[y]
			var incr int
			if !strings.EqualFold(string(str1[y-1]), string(str2[x-1])) {
				incr = 1
			}

			column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}
	return column[s1len]
}

func minimum(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
