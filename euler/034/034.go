/*
Project Euler problem number 34 (https://projecteuler.net/problem=34)

Digit factorials

145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: As 1! = 1 and 2! = 2 are not sums they are not included.

Solution By Peter Welsh
09/20/2022
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	const limit = 100000
	answer := getAnswer(limit)
	fmt.Println(answer) // 40730
}

func getAnswer(limit int) int {
	sumOfSums := 0
	factorialCache := make(map[int]int) // cache for digits' factorials
	sumCache := make(map[string]int)    // cache for digits' sums
	for i := 10; i <= limit; i++ {
		digits := getDigits(i)
		sum, exists := sumCache[digits]
		if !exists {
			sum = getSumOfDigitsFactorials(i, factorialCache)
			sumCache[digits] = sum
		}
		if i == sum {
			sumOfSums += i
		}
	}
	return sumOfSums
}

// Returns a string representation of the sorted digits of the positive integer n
// e.g. "123" given n=231
func getDigits(n int) string {
	var digits []int
	for n != 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	sort.Ints(digits)
	var result string
	for _, v := range digits {
		result += strconv.Itoa(v)
	}
	return result
}

// Returns the sum of the factorials of the digits of n
// e.g. 1! + 2! + 3! = 9 for n=123
func getSumOfDigitsFactorials(n int, cache map[int]int) int {
	sum := 0
	for n != 0 {
		sum += factorial(n%10, cache)
		n /= 10
	}
	return sum
}

// Returns the factorial of n
// e.g. 4! = 4*3*2*1 = 24
func factorial(n int, cache map[int]int) int {
	if result, exists := cache[n]; exists {
		return result
	}
	if n <= 1 {
		return 1
	}
	result := n * factorial(n-1, cache)
	cache[n] = result
	return result
}
