/*
Project Euler problem number 46 (https://projecteuler.net/problem=46)

Goldbach's other conjecture

It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.

9 = 7 + 2×1²
15 = 7 + 2×2²
21 = 3 + 2×3²
25 = 7 + 2×3²
27 = 19 + 2×2²
33 = 31 + 2×1²

It turns out that the conjecture was false.

What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?

Solution By Peter Welsh
09/27/2022
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	const firstOddComposite = 9
	start := firstOddComposite
	answer := getAnswer(start)
	fmt.Println(answer) // 5777
}

func getAnswer(start int) int {
	var answer, n int
	for answer == 0 {
		n++
		end := int(math.Pow10(n))
		answer = getFirstCompositeNumber(start, end)
		start = end + 1
	}
	return answer
}

func getFirstCompositeNumber(start, end int) int {
	cache := make(map[int]bool)
	for i := start; i < end; i += 2 {
		if isPrime(i, cache) {
			// skip this value; we are looking for a composite (non-prime) number
			continue
		}
		var n, x int
		for {
			n++
			x = 2 * n * n
			if x >= i {
				// i cannot be written as the sum of a prime and 2x a square
				// We have our answer!
				return i
			}
			if isPrime(i-x, cache) {
				// i can be written as the sum of a prime and 2x a square
				// The search continues...
				break
			}
		}
	}
	return 0
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
