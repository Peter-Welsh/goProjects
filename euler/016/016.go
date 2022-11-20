/*
Project Euler problem number 16 (https://projecteuler.net/problem=16)

Power digit sum

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?

Solution By Peter Welsh
09/12/2022
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	num := pow(2, 1000)
	answer := sumDigitsInt(num)
	fmt.Println(answer) //1366
}

func pow(x int64, y int64) *big.Int {
	num := big.NewInt(x)
	return num.Exp(num, big.NewInt(y), nil)
}

func sumDigitsInt(num *big.Int) int {
	return sumDigits(fmt.Sprintf("%v", num))
}

func sumDigits(num string) (sum int) {
	for _, digit := range num {
		sum += int(digit - '0')
	}
	return
}
