/*
Project Euler problem number 23 (https://projecteuler.net/problem=23)

Non-abundant sums

A perfect number is a number for which the sum of its proper divisors is exactly equal to the number.
For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24.
By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers.
However, this upper limit cannot be reduced any further by analysis even though it is known that the
greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.

Solution By Peter Welsh
09/14/2022
*/

package main

import (
	"fmt"
)

const limit = 28123

func main() {
	abundantNums := getAbundantNumsUpTo(limit)
	sum := 0
	for n := 1; n <= limit; n++ {
		qualifies := qualifies(n, abundantNums)
		if qualifies {
			sum += n
		}
	}
	fmt.Println(sum) // 4179871
}

// a number N qualifies if it cannot be written as the sum of two abundant numbers
func qualifies(n int, abundantNums []int) bool {
	left := 0
	right := len(abundantNums) - 1
	var sum int
	// using <= rather than < because we can use the same number twice (e.g. 12 and 12 to get 24)
	for left <= right {
		sum = abundantNums[left] + abundantNums[right]
		if sum > n {
			right--
		} else if sum < n {
			left++
		} else {
			return false
		}
	}
	return true
}

func getAbundantNumsUpTo(limit int) []int {
	var abundantNums []int
	var sum uint64
	var abundant bool
	for i := 1; i <= limit; i++ {
		sum = getSumOfDivisors(i)
		abundant = sum > uint64(i)
		if abundant {
			abundantNums = append(abundantNums, i)
		}
	}
	return abundantNums
}

func getSumOfDivisors(n int) uint64 {
	sum := uint64(0)
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += uint64(i)
		}
	}
	return sum
}
