/*
Project Euler problem number 7 (https://projecteuler.net/problem=7)

10001st prime

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

Solution By Peter Welsh
09/09/2022
*/

package main

import (
	"fmt"
)

func main() {
	n := 10001
	prime := findNthPrime(n)
	fmt.Printf("The %dth prime number is %d.", n, prime) //for n=10001, 104743
}

func findNthPrime(n int) int {
	if n == 1 {
		return 2
	}
	if n == 2 {
		return 3
	}
	prime := 3
	i := 2
	for i < n {
		for {
			prime += 2
			if isPrime(prime) {
				break
			}
		}
		i++
	}
	return prime
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n%2 == 0 || n%3 == 0 {
		return n == 2 || n == 3
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
