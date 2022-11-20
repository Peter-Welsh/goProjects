/*
Project Euler problem number 14 (https://projecteuler.net/problem=14)

Longest Collatz sequence

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.

Solution By Peter Welsh
09/12/2022
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	longest := 0
	var answer int
	for i := 1; i < 1000000; i++ {
		length := getCollatzLength(i)
		longest = int(math.Max(float64(longest), float64(length)))
		if longest == length {
			answer = i
		}
	}
	fmt.Println(answer) // 837799 (with a chain length of 524)
}

func getCollatzLength(n int) (collatzLength int) {
	collatz := n
	for {
		collatz = getCollatz(collatz)
		collatzLength++
		if collatz == 1 {
			break
		}
	}
	return
}

func getCollatz(n int) (collatz int) {
	if n%2 == 0 {
		collatz = n / 2
	} else {
		collatz = 3*n + 1
	}
	return
}
