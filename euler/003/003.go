/*
Project Euler problem number 3 (https://projecteuler.net/problem=3)

Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

Solution By Peter Welsh
09/08/2022
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	factor := getLargestPrimeFactor(600851475143)
	fmt.Println(factor) //6857
}

func getLargestPrimeFactor(n int) int {
	var largestPrimeFactor int
	for n%2 == 0 {
		n /= 2
		largestPrimeFactor = 2
	}

	for i := 3; float64(i) <= math.Sqrt(float64(n)); i += 2 {
		for n%i == 0 {
			largestPrimeFactor = i
			n = n / i
		}
	}

	if n > 2 {
		return n
	}
	return largestPrimeFactor
}
