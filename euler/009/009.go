/*
Project Euler problem number 9 (https://projecteuler.net/problem=9)

Special Pythagorean triplet

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2

For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

Solution By Peter Welsh
09/09/2022
*/

package main

import "fmt"

func main() {
	target := 1000
	a, b, c := getTriplet(target)
	fmt.Printf("The triplet that adds up to %d is (%d, %d, %d).\r\n", target, a, b, c)
	fmt.Printf("The product is %d.", a*b*c)
}

func getTriplet(target int) (int, int, int) {
	for a := 1; a < target; a++ {
		for b := a + 1; b < target-a; b++ {
			c := target - b - a
			if a*a+b*b == c*c {
				return a, b, c
			}
		}
	}
	return 0, 0, 0
}
