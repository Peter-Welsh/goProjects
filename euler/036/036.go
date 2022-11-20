/*
Project Euler problem number 36 (https://projecteuler.net/problem=36)

Double-base palindromes

The decimal number, 585 = 1001001001 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)

Solution By Peter Welsh
09/21/2022
*/

package main

import (
	"fmt"
)

func main() {
	const limit = 1000000
	answer := getAnswer(limit)
	fmt.Println(answer) // 872187
}

func getAnswer(limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		if isPalindrome(fmt.Sprintf("%d", i)) && isPalindrome(fmt.Sprintf("%b", i)) {
			sum += i
		}
	}
	return sum
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
