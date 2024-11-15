/*
Project Euler problem number 45 (https://projecteuler.net/problem=45)

Triangular, pentagonal, and hexagonal

Triangle, pentagonal, and hexagonal numbers are generated by the following formulae:
Triangle 	  	Tn=n(n+1)/2 	  	1, 3, 6, 10, 15, ...
Pentagonal 	  	Pn=n(3n−1)/2 	  	1, 5, 12, 22, 35, ...
Hexagonal 	  	Hn=n(2n−1) 	  	1, 6, 15, 28, 45, ...

It can be verified that T285 = P165 = H143 = 40755.

Find the next triangle number that is also pentagonal and hexagonal.

Solution By Peter Welsh
09/27/2022
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	const threshold = 40755
	answer := getAnswer(threshold)
	fmt.Println(answer) // 1533776805
}

func getAnswer(threshold int) int {
	var answer int
	var pow int
	for answer == 0 {
		pow++
		n := int(math.Pow10(pow))
		triags := getTriagonals(n)
		pentags := getPentagonals(n)
		hexags := getHexagonals(n)
		answer = getTriPenHexagonal(threshold, triags, pentags, hexags)
	}
	return answer
}

func getTriPenHexagonal(threshold int, triags, pentags, hexags map[int]struct{}) int {
	for t := range triags {
		if t <= threshold {
			continue
		}
		if _, exists := pentags[t]; !exists {
			continue
		}
		if _, exists := hexags[t]; exists {
			return t
		}
	}
	return 0
}

var Dummy struct{}

func getTriagonals(n int) map[int]struct{} {
	triags := make(map[int]struct{})
	for i := 1; i <= n; i++ {
		triags[i*(i+1)/2] = Dummy
	}
	return triags
}

func getPentagonals(n int) map[int]struct{} {
	pentags := make(map[int]struct{})
	for i := 1; i <= n; i++ {
		pentags[i*(3*i-1)/2] = Dummy
	}
	return pentags
}

func getHexagonals(n int) map[int]struct{} {
	hexags := make(map[int]struct{})
	for i := 1; i <= n; i++ {
		hexags[i*(2*i-1)] = Dummy
	}
	return hexags
}
