/*
Project Euler problem number 42 (https://projecteuler.net/problem=42)

Coded triangle numbers

The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value.
For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words, how many are triangle words?

Solution By Peter Welsh
09/23/2022
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	numTriangleWords := getAnswer()
	fmt.Print(numTriangleWords) // 162
}

func getAnswer() int {
	content, err := os.ReadFile("042/p042_words.txt")
	if err != nil {
		fmt.Println("Err")
	}
	var numTriangleWords int
	const alphabetSize = 26
	triangleNumbers := getTriangleNumbersUpTo(alphabetSize)
	// content is expected to look like: "WORD1", "WORD2", ...
	for _, v := range strings.Split(string(content), ",") {
		word := strings.Trim(v, "\"")
		if isTriangleWord(word, triangleNumbers) {
			numTriangleWords++
		}
	}
	return numTriangleWords
}

var Dummy struct{}

func getTriangleNumbersUpTo(n int) map[int]struct{} {
	var triangleNum int
	triangleNums := make(map[int]struct{})
	for i := 1; i <= n; i++ {
		triangleNum = (i * (i + 1)) / 2
		triangleNums[triangleNum] = Dummy
	}
	return triangleNums
}

func isTriangleWord(word string, triangleNums map[int]struct{}) bool {
	value := getValue(word)
	_, exists := triangleNums[value]
	return exists
}

func getValue(word string) int {
	value := 0
	for _, v := range word {
		// 'A' has a value of 1, 'B' is 2, etc.
		value += int(v-'A') + 1
	}
	return value
}
