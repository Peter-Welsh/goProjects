/*
Project Euler problem number 21 (https://projecteuler.net/problem=21)

Amicable numbers

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.

Solution By Peter Welsh
09/14/2022
*/

package main

import (
	"fmt"
)

func main() {
	sumOfSums := uint64(0)
	prevSum1 := uint64(0)
	prevSum2 := uint64(0)
	for n := uint64(1); n < uint64(10000); n++ {
		sum1 := getSumOfDivisors(n)
		sum2 := getSumOfDivisors(sum1)
		areAmicable := n == sum2 && n != sum1
		notAlreadyCounted := sum1 != prevSum2 || sum2 != prevSum1
		if areAmicable && notAlreadyCounted {
			sumOfSums += sum1 + sum2
			prevSum1 = sum1
			prevSum2 = sum2
		}
	}
	answer := sumOfSums
	fmt.Println(answer) // 31626
}

func getSumOfDivisors(n uint64) uint64 {
	sum := uint64(0)
	for i := uint64(1); i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
