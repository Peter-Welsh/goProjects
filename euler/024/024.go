/*
Project Euler problem number 24 (https://projecteuler.net/problem=24)

Lexicographic permutations

A permutation is an ordered arrangement of objects.
For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
If all of the permutations are listed numerically or alphabetically, we call it lexicographic order.
The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?

Solution By Peter Welsh
09/14/2022
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 1000000
	digits := "0123456789"
	answer := getAnswer(n, digits)
	fmt.Println(answer) // 2783915460
}

func getAnswer(n int, chars string) string {
	result := make([]int, 10)
	n--
	lastIndex := len(chars) - 1
	permNum := getFactorial(lastIndex)
	for i := lastIndex; i > 0; i-- {
		result[i] = n / permNum
		n %= permNum
		permNum /= i
	}
	// Result is [0 1 1 2 1 5 2 6 6 2] at this point for n = 1000000 and chars = "0123456789"
	// This number (2662512110) is the representation of 1000000 in the factorial number system
	// https://en.wikipedia.org/wiki/Factorial_number_system

	for i := 1; i <= lastIndex; i++ {
		for j := i - 1; j >= 0; j-- {
			if result[j] >= result[i] {
				result[j]++
			}
		}
	}
	answer := ""
	for i := 0; i <= lastIndex; i++ {
		answer = strconv.Itoa(result[i]) + answer
	}
	return answer
}

func getFactorial(n int) int {
	if n < 2 {
		return 1
	}
	return n * getFactorial(n-1)
}
