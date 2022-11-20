/*
Project Euler problem number 26 (https://projecteuler.net/problem=26)

Reciprocal cycles

A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

    1/2	= 	0.5
    1/3	= 	0.(3)
    1/4	= 	0.25
    1/5	= 	0.2
    1/6	= 	0.1(6)
    1/7	= 	0.(142857)
    1/8	= 	0.125
    1/9	= 	0.(1)
    1/10	= 	0.1

Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.

Solution By Peter Welsh
09/15/2022
*/

package main

import (
	"fmt"
)

func main() {
	answer := getDenominatorWithLongestCycle(1000)
	fmt.Println(answer) //983
}

func getDenominatorWithLongestCycle(limit int) int {
	var cycleLength int
	denominator := 2
	maxLength := 0
	for d := limit; d > 2; d-- {
		if d%2 != 0 && d%5 != 0 {
			cycleLength = getCycleLength(d)
			if cycleLength > maxLength {
				denominator = d
				maxLength = cycleLength
			}
		}
	}

	return denominator
}

func getCycleLength(number int) int {
	count := 1
	remainder := 10 % number

	for remainder != 1 {
		remainder = remainder * 10 % number
		count++
	}

	return count
}
