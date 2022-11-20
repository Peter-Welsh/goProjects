/*
Project Euler problem number 1 (https://projecteuler.net/problem=1)

Multiples of 3 or 5

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.

Solution By Peter Welsh
09/08/2022
*/

package main

import (
	"fmt"
)

func main() {
	sum := 0
	for v := 1; v < 1000; v++ {
		if v%3 == 0 || v%5 == 0 {
			sum += v
		}
	}
	fmt.Println(sum) //233168
}
