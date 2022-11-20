/*
Project Euler problem number 38 (https://projecteuler.net/problem=38)

Pandigital multiples

Take the number 192 and multiply it by each of 1, 2, and 3:

    192 × 1 = 192
    192 × 2 = 384
    192 × 3 = 576

By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)

The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).

What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?

Solution By Peter Welsh
09/22/2022
*/

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	answer := getAnswer()
	fmt.Println(answer) // 932718654 (formed by n=2 and integer=9327)
}

func getAnswer() int {
	var prod string
	largestProduct := 0
	for i := 1; i < 10000; i++ {
		for n := 1; n < 10; n++ {
			prod = getConcatenatedProduct(i, n)
			if isPandigital(prod) {
				product, _ := strconv.Atoi(prod)
				largestProduct = int(math.Max(float64(product), float64(largestProduct)))
			}
		}
	}
	return largestProduct
}

func getConcatenatedProduct(num int, n int) string {
	var result string
	for i := 1; i <= n; i++ {
		result += strconv.Itoa(num * i)
	}
	return result
}

func isPandigital(digits string) bool {
	oneToNine := "123456789"
	prevDigits := digits
	for _, v := range oneToNine {
		digits = strings.Replace(digits, fmt.Sprintf("%c", v), "", -1)
		// should have deleted 1 character if it is pandigital
		if len(prevDigits)-len(digits) != 1 {
			return false
		}
		prevDigits = digits
	}
	return len(digits) == 0 // should be an empty string if it is pandigital
}
