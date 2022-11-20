/*
Project Euler problem number 28 (https://projecteuler.net/problem=28)

Number spiral diagonals

Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:

21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13

It can be verified that the sum of the numbers on the diagonals is 101.

What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?

Solution By Peter Welsh
09/16/2022
*/

package main

import (
	"fmt"
)

func main() {
	n := 1001
	diagonalSum := getDiagonalSum(n)
	fmt.Println(diagonalSum) // 669171001
}

func getDiagonalSum(n int) int {
	m := 3           // dimensions of current spiral
	diagonalSum := 1 // center of spiral
	onDiagonal := true
	for i := 2; i <= n*n; i++ {
		onDiagonal = i%(m-1) == 1
		if onDiagonal {
			diagonalSum += i
		}
		if m*m == i {
			m += 2
		}
	}
	return diagonalSum
}
