/*
Project Euler problem number 37 (https://projecteuler.net/problem=37)

Truncatable primes

The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right,
and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.

Solution By Peter Welsh
09/21/2022
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	const start = 11 // first two-digit prime
	const end = 999999
	answer := getAnswer(start, end)
	fmt.Println(answer) // 748317
}

func getAnswer(start, end int) int {
	sum := 0
	// cache for prime numbers to avoid checking the same number twice
	cache := make(map[int]bool)
	// declare modifiers once, outside the loop
	modifiers := []func(*int){removeFirstDigit, removeLastDigit}
	// increment n by 6 each iteration and check n and n+2 (common method of prime-checking optimization)
	for n := start; n < end; n += 6 {
		if isLeftRightTruncatablePrime(n, cache, modifiers) {
			sum += n
		}
		if isLeftRightTruncatablePrime(n+2, cache, modifiers) {
			sum += (n + 2)
		}
	}
	return sum
}

// Checks if n is both a left-truncatable prime and a right-truncatable prime
// cache is a collection of numbers that have been already checked for primeness
// modifiers are two functions that either remove the first digit or the last digit of n at each iteration
func isLeftRightTruncatablePrime(n int, cache map[int]bool, modifiers []func(*int)) bool {
	for _, modifier := range modifiers {
		if !isTruncatablePrime(n, cache, modifier) {
			return false
		}
	}
	return true
}

func isTruncatablePrime(n int, cache map[int]bool, modifier func(*int)) bool {
	for n > 0 {
		if !isPrime(n, cache) {
			return false
		}
		modifier(&n)
	}
	return true
}

func removeLastDigit(n *int) {
	*n /= 10
}

func removeFirstDigit(n *int) {
	*n %= int(math.Pow10(len(strconv.Itoa(*n)) - 1))
}

func isPrime(n int, cache map[int]bool) bool {
	if isPrime, exists := cache[n]; exists {
		return isPrime
	}
	if n <= 1 {
		// all primes are greater than 1
		cache[n] = false
		return false
	}
	if n%2 == 0 || n%3 == 0 {
		// 2 is a prime and 3 is a prime
		// but all other numbers evenly divisible by 2 or 3 are not prime
		cache[n] = n == 2 || n == 3
		return cache[n]
	}
	for i := 5; i <= int(math.Sqrt(float64(n))); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			cache[n] = false
			return false
		}
	}
	cache[n] = true
	return true
}
