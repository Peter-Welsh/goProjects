/*
Project Euler problem number 30 (https://projecteuler.net/problem=30)

Digit fifth powers

Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

    1634 = 1^4 + 6^4 + 3^4 + 4^4
    8208 = 8^4 + 2^4 + 0^4 + 8^4
    9474 = 9^4 + 4^4 + 7^4 + 4^4

As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.

Solution By Peter Welsh
09/16/2022
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	const pow = 5
	const start = 2
	end := getUpperBounds(pow)
	sum := 0
	for i := start; i < end; i++ {
		if qualifies(i, pow) {
			sum += i
		}
	}
	fmt.Println(sum) //443839 for pow = 5
}

func getUpperBounds(pow int) int {
	const highestDigit = 9
	x := int(math.Pow(float64(highestDigit), float64(pow))) // 9 ^ 5
	return len(strconv.Itoa(x)) * x                         // 6 * (9 ^ 5)
}

// a number qualifies if the sum of its digits each raised to the given power equals that number
// e.g. for pow = 5, num 4150 qualifies because 4^5 + 1^5 + 5^5 + 0^5 = 4150
func qualifies(num int, pow int) bool {
	n := num
	sum := 0
	for n > 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(pow)))
		n /= 10
	}
	return sum == num
}
