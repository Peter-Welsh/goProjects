/*
Project Euler problem number 35 (https://projecteuler.net/problem=35)

Circular primes

The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.

There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.

How many circular primes are there below one million?

Solution By Peter Welsh
09/20/2022
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	const threshold = 1000000
	answer := getAnswer(threshold)
	fmt.Println(answer) // 55
}

func getAnswer(threshold int) int {
	cache := make(map[int]bool)
	count := 1 // 2 is the first prime
	// start at 3 so we can skip even numbers
	for i := 3; i < threshold; i += 2 {
		if isCircularPrime(i, cache) {
			count++
		}
	}
	return count
}

func isCircularPrime(n int, cache map[int]bool) bool {
	// 357 -> 735 -> 573
	x := len(strconv.Itoa(n))
	for i := 0; i < x; i++ {
		tmp := n % 10
		n /= 10
		n += tmp * int(math.Pow(10, float64(x-1)))
		if !isPrime(n, cache) {
			return false
		}
	}
	return true
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
