/*
Project Euler problem number 10 (https://projecteuler.net/problem=10)

Summation of primes

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.

Solution By Peter Welsh
09/09/2022
*/

package main

import (
	"fmt"
	"math"
)

const threshold = 2000000

func main() {
	sum := getSumOfPrimesUpTo(threshold)
	// 142913828922 for threshold=2000000
	fmt.Printf("The sum of all prime numbers from 1 to %d is %d.", threshold, sum)
}

func getSumOfPrimesUpTo(threshold int) int {
	sum := 0
	sum += 2
	sum += 3
	for i := 5; i < threshold; i += 6 {
		if isPrime(i) {
			sum += i
		}
		if isPrime(i + 2) {
			sum += i + 2
		}
	}
	return sum
}

func isPrime(n int) bool {
	if n == 1 {
		// 1 is not a prime number
		return false
	}
	if n%2 == 0 || n%3 == 0 {
		// 2 is a prime and 3 is a prime
		return n == 2 || n == 3
	}
	for i := 5; i <= int(math.Sqrt(float64(n))); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
