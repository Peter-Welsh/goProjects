/*
Project Euler problem number 50 (https://projecteuler.net/problem=50)

Consecutive prime sum

The prime 41, can be written as the sum of six consecutive primes:
41 = 2 + 3 + 5 + 7 + 11 + 13

This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?

Solution By Peter Welsh
09/28/2022
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	const target = 1000000
	answer := getAnswer(target)
	fmt.Println(answer) // 997651 (sum of the 543 consecutive primes starting with 7 and ending with 3931)
}

// this technique works fine for target=1000 and target=1000000 but doesn't work for target=100
func getAnswer(target int) int {
	cache := make(map[int]bool)
	var sum int
	var primes []int
	for i := 0; sum+i < target; i++ {
		if !isPrime(i, cache) {
			continue
		}
		primes = append(primes, i)
		sum += i
	}
	// remove primes from the start of the sequence until the sum is prime
	for i := 0; !isPrime(sum, cache); i++ {
		sum -= primes[i]
	}
	return sum
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
