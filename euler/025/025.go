/*
Project Euler problem number 25 (https://projecteuler.net/problem=25)

1000-digit Fibonacci number

The Fibonacci sequence is defined by the recurrence relation:

    Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.

Hence the first 12 terms will be:

    F1 = 1
    F2 = 1
    F3 = 2
    F4 = 3
    F5 = 5
    F6 = 8
    F7 = 13
    F8 = 21
    F9 = 34
    F10 = 55
    F11 = 89
    F12 = 144

The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

Solution By Peter Welsh
09/15/2022
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	numDigits := int64(1000)
	threshold := *pow(10, numDigits-1) // 10^999 has 1000 digits
	index := getIndexOfFirstFibonacciTermGreaterThanOrEqualTo(&threshold)
	fmt.Println(index) // 4782
}

func getIndexOfFirstFibonacciTermGreaterThanOrEqualTo(threshold *big.Int) int {
	prevTerm := big.NewInt(1)
	prevPrev := big.NewInt(1)
	index := 2
	result := big.NewInt(0)
	for result.Cmp(threshold) == -1 {
		result.Add(prevTerm, prevPrev)
		index++
		prevTerm.Set(prevPrev)
		prevPrev.Set(result)
	}
	return index
}

func pow(x int64, y int64) *big.Int {
	num := big.NewInt(x)
	return num.Exp(num, big.NewInt(y), nil)
}
