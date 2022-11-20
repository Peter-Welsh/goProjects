/*
Project Euler problem number 41 (https://projecteuler.net/problem=41)

Pandigital prime

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once. For example, 2143 is a 4-digit pandigital and is also prime.

What is the largest n-digit pandigital prime that exists?

Solution By Peter Welsh
09/23/2022
*/

package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	answer := getAnswer()
	fmt.Println(answer) // 7652413
}

func getAnswer() int {
	// None of the 9- or 8-digit pandigital numbers are prime.
	// This can be proven by looking at the sum of their digits, which is constant.
	// The sum of the digits are 45 and 36 respectively for 9- and 8-digit pandigital numbers, both of which are evenly divisible by 3.
	// Given that the sum of their digits is divisible by 3, we know that the number itself is evenly divisible by 3 and thus not prime.
	// The sum of the digits of all 7-digit pandigital numbers is 28, so our answer has 7 digits.
	const n = 7
	start := getStart(n)
	end := getEnd(n)
	digits := getDigits(n)
	for i := start; i >= end; i -= 2 {
		if isPandigital(i, digits) && isPrime(i) {
			return i
		}
	}
	return 0
}

// Returns a string of the digits from 1 to n
// e.g. "1234567" for n=7
func getDigits(n int) string {
	var digits string
	for i := 1; i <= n; i++ {
		digits += strconv.Itoa(i)
	}
	return digits
}

// Returns the highest n-digit pandigital number
// e.g. 7654321 for n=7
func getStart(n int) int {
	result := 0
	for i := n; i > 0; i-- {
		result += int(math.Pow10(i-1)) * i
	}
	return result
}

// Returns the smallest n-digit pandigital number
// e.g. 1234567 for n=7
func getEnd(n int) int {
	result := 0
	originalN := n
	for i := 1; i <= originalN; i++ {
		result += int(math.Pow10(n-1)) * i
		n--
	}
	return result
}

func isPrime(n int) bool {
	if n <= 1 {
		// all primes are greater than 1
		return false
	}
	if n%2 == 0 || n%3 == 0 {
		// 2 is a prime and 3 is a prime
		// but all other numbers evenly divisible by 2 or 3 are not prime
		return n == 2 || n == 3
	}
	for i := 5; i <= int(math.Sqrt(float64(n))); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Checks if a number is pandigital by sorting its string representation in ascending order and checking for equality against the digits passed in.
// Assumes that digits are sorted in ascending order, e.g. "1234567"
func isPandigital(num int, digits string) bool {
	return digits == SortString(strconv.Itoa(num))
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
