/*
Project Euler problem number 32 (https://projecteuler.net/problem=32)

Pandigital products

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.

Solution By Peter Welsh
09/16/2022
*/

package main

import (
	"fmt"
)

// in order to use up 9 digits, each number in the equation a * b = c cannot be greater than 9999
const maxNumber = 9999

func main() {
	sumOfProducts := getSumOfProducts()
	fmt.Println(sumOfProducts) // 45228
}

var Dummy struct{}

func getSumOfProducts() int {
	sum := 0
	var product int
	seen := make(map[int]struct{})
	for multiplicand := 1; multiplicand < maxNumber; multiplicand++ {
		for multiplier := 1; multiplier < maxNumber; multiplier++ {
			product = multiplicand * multiplier
			if product > maxNumber {
				break
			}
			if qualifies(multiplicand, multiplier, product) {
				if _, exists := seen[product]; exists {
					// skip the duplicate
					continue
				}
				seen[product] = Dummy
				sum += product
			}
		}
	}
	return sum
}

// three numbers qualify if the digits 1 thru 9 appear exactly once
// e.g. the three numbers 1963, 4, and 7852 qualify
func qualifies(x int, y int, z int) bool {
	digits := fmt.Sprintf("%d%d%d", x, y, z)
	return isPandigital(digits)
}

func isPandigital(digits string) bool {
	seen := make(map[rune]struct{})
	if len(digits) != 9 {
		return false
	}
	for _, digit := range digits {
		if digit == '0' {
			return false
		}
		if _, exists := seen[digit]; exists {
			return false
		} else {
			seen[digit] = Dummy
		}
	}
	return len(seen) == 9
}
