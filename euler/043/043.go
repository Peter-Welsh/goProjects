/*
Project Euler problem number 43 (https://projecteuler.net/problem=43)

Sub-string divisibility

The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order,
but it also has a rather interesting sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:

    d2d3d4=406 is divisible by 2
    d3d4d5=063 is divisible by 3
    d4d5d6=635 is divisible by 5
    d5d6d7=357 is divisible by 7
    d6d7d8=572 is divisible by 11
    d7d8d9=728 is divisible by 13
    d8d9d10=289 is divisible by 17

Find the sum of all 0 to 9 pandigital numbers with this property.

Solution By Peter Welsh
09/26/2022
*/

package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	const start = 1023456789
	const end = 9876543210
	sum := getAnswer(start, end)
	fmt.Println(sum) // 16695334890
}

func getAnswer(start, end int) int {
	sum := 0
	digits := getDigits(9)
	var inc int
	flag := true
	for i := start; i <= end; i += inc {
		if matchesPattern(i) && isPandigital(i, digits) {
			sum += i
		}
		last6digits := i % int(math.Pow10(6))
		// optimization: all the numbers we want end with 952867 or 357289 due to the constraints given in the problem description
		if flag {
			inc = 952867 - last6digits
		} else {
			inc = 357289 - last6digits + int(math.Pow10(6))
		}
		flag = !flag
	}
	return sum
}

func matchesPattern(n int) bool {
	str := strconv.Itoa(n)
	primes := []int{2, 3, 5, 7, 11, 13, 17}
	start := len(primes)
	end := len(str)
	for start > 0 {
		var slice, _ = strconv.Atoi(str[start:end])
		if slice%primes[start-1] != 0 {
			return false
		}
		start--
		end--
	}
	return true
}

// Returns a string of the digits from 0 to n
// e.g. "01234567" for n=7
func getDigits(n int) string {
	var digits string
	for i := 0; i <= n; i++ {
		digits += strconv.Itoa(i)
	}
	return digits
}

// Checks if a number is pandigital by sorting its string representation in ascending order and checking for equality against the digits passed in.
// Assumes that digits are sorted in ascending order, e.g. "01234567"
func isPandigital(num int, digits string) bool {
	return digits == SortString(strconv.Itoa(num))
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
