/*
Project Euler problem number 39 (https://projecteuler.net/problem=39)

Integer right triangles

If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.

{20,48,52}, {24,45,51}, {30,40,50}

For which value of p ≤ 1000, is the number of solutions maximised?

Solution By Peter Welsh
09/22/2022
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	const limit = 1000
	answer := getAnswer(limit)
	fmt.Println(answer) // 840
}

func getAnswer(limit int) int {
	var solutions, answer, max int
	const numSides = 3
	for p := numSides; p <= limit; p++ {
		solutions = 0
		for a := 1; a < p/numSides; a++ {
			// a, b, and c need to satisfy: 1) a+b>=c 2) a<b<c 3) a+b+c == p
			for b := int(math.Max(float64(p/2-a), float64(a))); b < p-b-a; b++ {
				c := p - b - a
				if a*a+b*b == c*c {
					solutions++
				}
			}
		}
		if solutions > max {
			max = solutions
			answer = p
		}
	}
	return answer
}
