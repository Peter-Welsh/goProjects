/*
This program prompts the user for a number and determines whether it is prime.
If it is not prime, the program displays the number's factors.

Limitations:
- the program is very slow for some large numbers, e.g. 1591591591491
- the input type is int64, which has a maximum value of 9223372036854775807

By Peter Welsh
08-22-2022
*/

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	for {
		number := getNumberInput()
		isPrime := isPrime(number)
		var factors []int64
		if !isPrime {
			factors = getFactors(number)
		}

		printResult(isPrime, number, factors)

		time.Sleep(time.Second)
		var continueOrQuit string
		fmt.Println("\r\nPress enter to restart or Q to quit.")
		fmt.Scanln(&continueOrQuit)
		if strings.EqualFold(continueOrQuit, "Q") {
			break
		}
	}
}

func getNumberInput() int64 {
	var number int64
	for {
		fmt.Print("Enter a positive integer: ")
		fmt.Scanln(&number)
		if number > 0 {
			break
		}
		fmt.Print("Try again. ")
	}
	return number
}

func isPrime(n int64) bool {
	if n == 1 {
		return false
	}

	if n%2 == 0 || n%3 == 0 {
		return n == 2 || n == 3
	}

	for i := int64(5); i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func getFactors(n int64) []int64 {
	factors := []int64{}
	c := int64(2)
	for n > 1 {
		if n%c == 0 {
			n /= c
			factors = append(factors, c)
		} else {
			c++
		}
	}
	return factors
}

func printResult(isPrime bool, number int64, factors []int64) {
	if isPrime {
		fmt.Printf("%v is a prime number!\r\n", number)
	} else {
		fmt.Printf("%v is NOT a prime number.\r\n", number)
		fmt.Println("Its factors are:", factors)
	}
}
