/*
Project Euler problem number 33 (https://projecteuler.net/problem=33)

Digit cancelling fractions

The fraction 49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify
it may incorrectly believe that 49/98 = 4/8, which is correct, is obtained by cancelling the 9s.

We shall consider fractions like, 30/50 = 3/5, to be trivial examples.

There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.

If the product of these four fractions is given in its lowest common terms, find the value of the denominator.

Solution By Peter Welsh
09/20/2022
*/

package main

import (
	"fmt"
)

const smallestTwoDigitNumber = 10
const highestTwoDigitNumber = 99

func main() {
	answer := getAnswer()
	fmt.Println(answer) // 100
}

func getAnswer() int {
	numeratorProduct := 1
	denominatorProduct := 1
	for i := smallestTwoDigitNumber; i < highestTwoDigitNumber; i++ {
		for j := i + 1; j <= highestTwoDigitNumber; j++ {
			if qualifies(i, j) {
				numeratorProduct *= i
				denominatorProduct *= j
			}
		}
	}
	/* 16/64, 19/95, 26/65, 49/98 */
	return denominatorProduct / numeratorProduct
}

func qualifies(numerator, denominator int) bool {
	fraction := float64(numerator) / float64(denominator)
	n2 := numerator % 10   //2nd digit in numerator
	n1 := numerator / 10   //1st "     "  "
	d2 := denominator % 10 //2nd digit in denominator
	d1 := denominator / 10 //1st "     "  "
	haveSharedDigit := (d1 == n1 || d1 == n2) || (d2 == n1 || d2 == n2)
	if !haveSharedDigit {
		return false
	}
	isTrivial := n2 == 0 && d2 == 0
	if isTrivial {
		return false
	}
	n1d1 := n1 == d1 && float64(n2)/float64(d2) == fraction
	n1d2 := n1 == d2 && float64(n2)/float64(d1) == fraction
	n2d1 := n2 == d1 && float64(n1)/float64(d2) == fraction // 16/64 -> 1/4 (sixes cancel)
	n2d2 := n2 == d2 && float64(n1)/float64(d1) == fraction
	cancelingGivesSameFraction := n1d1 || n1d2 || n2d1 || n2d2
	return cancelingGivesSameFraction
}
