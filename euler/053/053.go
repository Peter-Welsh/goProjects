/*
Project Euler problem number 53 (https://projecteuler.net/problem=53)

There are exactly ten ways of selecting three from five, 12345:
123, 124, 125, 134, 135, 145, 234, 235, 245, and 345

In combinatorics, we use the notation, "5 choose 3" = 10.

In general, "n choose r" = n!/(r!(n-r)!), where r <= n, n! = n * (n-1) * ... * 3 * 2 * 1, and 0! = 1.

It is not until n = 23, that a value exceeds one-million: "23 choose 10" = 1144066.

How many, not necessarily distinct, values of "n choose r" for 1 <= n <= 100, are greater than one-million?

Solution By Peter Welsh
09/29/2022
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	const threshold = 1000000
	const limit = 100
	answer := getAnswer(threshold, limit)
	fmt.Println(answer) // 4075
}

var Dummy = big.NewInt(1)

// Returns the number of values of (n r) that are greater than the threshold where n <= limit
func getAnswer(threshold, limit int) int {
	count := 0
	result := Dummy
	bottom := Dummy
	for n := 1; n <= limit; n++ {
		top := factorial(n)
		for r := n; r > 0; r-- {
			bottom.Mul(factorial(r), factorial(n-r))
			result.Div(top, bottom)
			if result.Cmp(big.NewInt(int64(threshold))) == 1 {
				count++
			}
		}
	}
	return count
}

// Returns the factorial of n
// e.g. 4! = 4*3*2*1 = 24
func factorial(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(1)
	}

	return big.NewInt(1).Mul(big.NewInt(int64(n)), factorial(n-1))
}
