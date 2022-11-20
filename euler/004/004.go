/*
Project Euler problem number 4 (https://projecteuler.net/problem=4)

Largest palindrome product

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.

Solution By Peter Welsh
09/09/2022
*/

package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	var n int
	flag.IntVar(&n, "n", 0, "Specify the number of digits N.")
	flag.Parse()
	if n < 2 || n > 10 {
		fmt.Println("N must be a positive integer between 2 and 10. Exiting.")
		fmt.Println("Usage: go run 004.go -n=3")
		return
	}
	var product, factor1, factor2 = getLargestPalindromicProduct(n)
	fmt.Printf("The largest palindromic product of two %d-digit integers is %d.", n, product)
	fmt.Println("")
	fmt.Printf("The factors are %d and %d.", factor1, factor2)

	//n=2 -> 9009 (99 and 91)
	//n=3 -> 906609 (993 and 913)
	//n=4 -> 99000099 (9999 and 9901)
	//n=5 -> 9966006699 (99979 and 99681)
	//n=6 -> 999000000999 (999999 and 999001)
	//n=7 -> 99956644665999 (9998017 and 9997647)
	//n=8 -> 9999000000009999 (99999999 and 99990001)
	//n=9 -> 999900665566009999 (999980347 and 999920317)
}

// The strategy is start at (10^n)-1 and work our way down, checking palindromic products until we find two factors that can produce it
func getLargestPalindromicProduct(n int) (uint64, uint64, uint64) {
	// for n=3, pow is 1000; start is 999; end is 900
	pow := uint64(math.Pow(10, float64(n)))
	start := pow - 1
	end := 9 * (pow / 10)
	for leftHalf := start; leftHalf >= end; leftHalf-- {
		// for n=3, palindromicProduct is 999999 on 1st iteration; 998899 on 2nd iteration...
		palindromicProduct := leftHalf*pow + reverse(leftHalf)

		// -=2 to skip over even numbers
		for factor1 := start; factor1 > end; factor1 -= 2 {
			if factor1%10 == 5 {
				// we can skip this candidate since the product needs to end with a 9
				// (if one factor ends with a 5, then the product will either end with a 0 or 5)
				continue
			}
			factor2 := palindromicProduct / factor1
			if factor2 > start {
				// we need an N-digit factor; this one has too many digits
				break
			}
			if palindromicProduct%factor1 == 0 {
				return palindromicProduct, factor1, factor2
			}
		}
	}
	return 0, 0, 0
}

func reverse(num uint64) uint64 {
	reversedNum := uint64(0)
	for num > 0 {
		reversedNum = num%10 + reversedNum*10
		num /= 10
	}
	return reversedNum
}
