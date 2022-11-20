/*
Project Euler problem number 52 (https://projecteuler.net/problem=52)

Permuted multiples

It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a different order.

Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.

Solution By Peter Welsh
09/29/2022
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	answer := getAnswer()
	fmt.Println(answer) // 142857
}

// 285714, 428571, 571428, 714285, 857142
func getAnswer() int {
	multiples := []int{2, 3, 4, 5, 6}
	for n := 1; ; n++ {
		isPerm := true
		x := strconv.Itoa(multiples[0] * n)
		for _, multiple := range multiples[1:] {
			y := strconv.Itoa(multiple * n)
			isPerm = isPerm && isPermutation(x, y)
		}
		if isPerm {
			return n
		}
	}
}

func isPermutation(x, y string) bool {
	return SortString(x) == SortString(y)
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
