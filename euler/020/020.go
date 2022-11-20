/*
Project Euler problem number 20 (https://projecteuler.net/problem=20)

Factorial digit sum

n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!

Solution By Peter Welsh
09/14/2022
*/

package main

import "fmt"

func main() {
	answer := getDigitSumOfFactorial(100)
	fmt.Println(answer) //648
}

// the trick is to use an array to store the digits of the big number
// could also use the "math/big" library to accomplish this
func getDigitSumOfFactorial(n int) int {
	factorial := getFactorial(n)
	digitSum := 0
	for i := 0; i < len(factorial); i++ {
		digitSum += factorial[i]
	}
	return digitSum
}

func getFactorial(n int) []int {
	var factorial []int
	factorial = append(factorial, 1)
	for num := 2; num <= n; num++ {
		multiply(&factorial, num)
	}
	return factorial
}

func multiply(num1 *[]int, num2 int) {
	carry := 0
	product := 0
	for i := 0; i < len(*num1); i++ {
		product = (*num1)[i]*num2 + carry
		(*num1)[i] = product % 10
		carry = product / 10
	}
	for carry > 0 {
		(*num1) = append((*num1), carry%10)
		carry /= 10
	}
}
