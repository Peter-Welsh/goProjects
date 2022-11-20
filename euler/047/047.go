/*
Project Euler problem number 47 (https://projecteuler.net/problem=47)

Distinct primes factors

The first two consecutive numbers to have two distinct prime factors are:

14 = 2 × 7
15 = 3 × 5

The first three consecutive numbers to have three distinct prime factors are:

644 = 2² × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.

Find the first four consecutive integers to have four distinct prime factors each. What is the first of these numbers?

Solution By Peter Welsh
09/27/2022
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	answer := getAnswer()
	fmt.Println(answer) // 134043
}

func getAnswer() int {
	const n = 4
	nth := 0
	cache := make(map[int]bool)
	previous := make([]bool, n)
	for i := 0; i < n-1; i++ {
		previous[i] = has4distinctPrimefactors(i+1, cache)
	}
	for i := n; ; i++ {
		if has4distinctPrimefactors(i, cache) {
			isNthInStreak := isStreak(n, previous)
			if isNthInStreak {
				nth = i
				break
			}
			shift(n, previous, true)
			continue
		}
		shift(n, previous, false)
	}
	return nth - n + 1
}

func isStreak(n int, previous []bool) bool {
	foundAnswer := false
	for j := 0; j < n-1; j++ {
		if !previous[j] {
			foundAnswer = false
			break
		} else {
			if j == n-2 {
				foundAnswer = true
			}
		}
	}
	return foundAnswer
}

func shift(n int, previous []bool, final bool) {
	for j := 0; j < n-1; j++ {
		previous[j] = previous[j+1]
	}
	previous[n-2] = final
}

// need to clean this up so that 4 is not hard-coded
func has4distinctPrimefactors(n int, cache map[int]bool) bool {
	numFactors := 0
	limit := int(math.Sqrt(float64(n)))
	prevN := n
	var firstPrime, secondPrime, thirdPrime int
	for i := 2; i <= limit; i++ {
		if !isPrime(i, cache) {
			continue
		}
		for n%i == 0 {
			n /= i
		}
		if prevN != n {
			if firstPrime == 0 {
				firstPrime = i
			} else if secondPrime == 0 {
				secondPrime = i
			} else if thirdPrime == 0 {
				thirdPrime = i
			}
			numFactors++
		}
		prevN = n
	}
	found4factors := numFactors == 4 && n == 1
	nIs4thFactor := numFactors == 3 && isPrime(n, cache)
	factorsAreDistinct := n != firstPrime && n != secondPrime && n != thirdPrime
	return (found4factors || nIs4thFactor) && factorsAreDistinct
}

func isPrime(n int, cache map[int]bool) bool {
	if isPrime, exists := cache[n]; exists {
		// check if the answer is cached
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
