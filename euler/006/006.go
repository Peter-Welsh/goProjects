/*
Project Euler problem number 6 (https://projecteuler.net/problem=6)

Sum square difference

The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 55^2 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 - 385 = 2640

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

Solution By Peter Welsh
09/09/2022
*/

package main

import (
	"fmt"
)

func main() {
	n := 100
	sumOfSquares := getSumOfSquares(n)
	squareOfSums := getSquareOfSums(n)
	difference := squareOfSums - sumOfSquares
	fmt.Printf("For the first %d natural numbers, the difference between the sum of squares and the square of sums is %d.", n, difference)
	// for n=100, 25164150
}

func getSumOfSquares(n int) int {
	var sum int
	for x := 1; x <= n; x++ {
		sum += (x * x)
	}
	return sum
}

func getSquareOfSums(n int) int {
	var sum int
	for x := 1; x <= n; x++ {
		sum += x
	}
	return sum * sum
}
