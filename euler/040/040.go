/*
Project Euler problem number 40 (https://projecteuler.net/problem=40)

Champernowne's constant

An irrational decimal fraction is created by concatenating the positive integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000

Solution By Peter Welsh
09/23/2022
*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	const limit = 1000000
	answer := getAnswer(limit)
	fmt.Println(answer) // 210
}

func getAnswer(limit int) int {
	builder := strings.Builder{}
	for i := 1; builder.Len() < limit; i++ {
		builder.WriteString(strconv.Itoa(i))
	}
	fraction := builder.String()
	product := 1
	for i := 0; float64(i) <= math.Log10(float64(limit)); i++ {
		digit := fraction[int(math.Pow10(i))-1]
		product *= int(digit - '0')
	}
	return product
}
