/*
Project Euler problem number 48 (https://projecteuler.net/problem=48)

Self powers

The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.

Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.

Solution By Peter Welsh
09/28/2022
*/

package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	const threshold = 1000
	answer := getAnswer(threshold)
	fmt.Println(answer) // 9110846700
}

func getAnswer(threshold int) (lastTenDigits string) {
	var result int64
	pow := big.NewInt(0)
	tenToTheTen := int64(math.Pow10(10))
	bigTenToTheTen := big.NewInt(tenToTheTen)
	for i := 1; i <= threshold; i++ {
		// set pow to i^i
		bigI := big.NewInt(int64(i))
		pow.Exp(bigI, bigI, nil)

		// add the last 10 digits to result
		lastTen := pow.Mod(pow, bigTenToTheTen)
		result += lastTen.Int64()

		// keep just the last ten digits
		result %= tenToTheTen
	}
	// format result to 10 characters, using leading zeros if needed
	return fmt.Sprintf("%010d", result)
}
