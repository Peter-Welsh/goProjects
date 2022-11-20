/*
Project Euler problem number 5 (https://projecteuler.net/problem=5)

Smallest multiple

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

Solution By Peter Welsh
09/09/2022
*/

package main

import (
	"fmt"
)

func main() {
	n := 20
	multiple := getSmallestMultiple(n)
	fmt.Println(multiple) //for n=20, 232792560
}

func getSmallestMultiple(n int) int {
	if n%10 != 0 || n < 1 {
		panic("Not implemented")
	}
	candidate := n
	const minFactor = 2
	// the answer must end with a 0 (be divisible by both 2 and 5)
	for {
		for x := n; x >= minFactor; x-- {
			if candidate%x == 0 {
				if x == minFactor {
					return candidate
				}
				continue
			}
			break
		}
		candidate += n
	}
}
