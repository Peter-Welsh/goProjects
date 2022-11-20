/*
Project Euler problem number 27 (https://projecteuler.net/problem=27)

Euler discovered the remarkable quadratic formula:
n^2 + n + 41
It turns out that the formula will produce 40 primes for the consecutive integer values 0 <= n <= 39.
However, when n = 40, 40^2 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and certainly when n = 41, 41^2 + 41 + 41 is clearly divisible by 41.
The incredible formula n^2 - 79n + 1601 was discovered, which produces 80 primes for the consecutive values 0 <= n <= 79.
The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

n^2 + an + b, where |a| < 1000 and |b| <= 1000 where |n| is the modulus/absolute value of n e.g. |11| = 11 and |-4| = 4

Find the product of the coefficients, a and b, for the quadratic expression that produces the maximum number of primes for consecutive values of n, starting with n = 0.

Solution By Peter Welsh
09/15/2022
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	limitA := 1000
	limitB := 1000
	a, b := getQuadraticCoefficients(limitA, limitB)
	answer := a * b     //-61 and 971
	fmt.Println(answer) //-59231
	/* the 71 consecutive prime numbers are:
	971, 911, 853, 797, 743, 691, 641, 593, 547, 503, 461, 421, 383, 347, 313, 281, 251, 223, 197,
	173, 151, 131, 113, 97, 83, 71, 61, 53, 47, 43, 41, 41, 43, 47, 53, 61, 71, 83, 97, 113, 131,
	151, 173, 197, 223, 251, 281, 313, 347, 383, 421, 461, 503, 547, 593, 641, 691, 743, 797, 853,
	911, 971, 1033, 1097, 1163, 1231, 1301, 1373, 1447, 1523, 1601 */
}

func getQuadraticCoefficients(limitA, limitB int) (int, int) {
	var c, numPrimes, maxPrimes int
	a := limitA - 1
	b := limitB
	finalA := 0
	finalB := 0
	for math.Abs(float64(a)) < float64(limitA) {
		for math.Abs(float64(b)) <= float64(limitB) {
			for n := 0; ; n++ {
				c = n*n + a*n + b
				if !isPrime(c) {
					break
				}
				numPrimes++
			}
			if maxPrimes < numPrimes {
				maxPrimes = numPrimes
				finalA = a
				finalB = b
			}
			numPrimes = 0
			b--
		}
		b = limitB // reset B
		a--
	}
	return finalA, finalB
}

func isPrime(n int) bool {
	if n <= 1 {
		// all primes are greater than 1
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
