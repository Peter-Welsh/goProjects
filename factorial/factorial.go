/*
This program prompts the user for a number and calculates the factorial.
The user can choose the method of computation, either recursive or iterative.

By Peter Welsh
08-22-2022
*/

package main

import (
	"fmt"
	"math/big"
	"strings"
	"time"
)

type Mode int8

const (
	Invalid Mode = iota
	Recursive
	Iterative
)

func main() {
	for {
		number := getNumberInput()
		mode := getModeInput()
		result := compute(number, mode)
		fmt.Printf("%v! = %v.", number, result)

		time.Sleep(time.Second)
		var continueOrQuit string
		fmt.Println("\r\n\r\nPress enter to restart or Q to quit.")
		fmt.Scanln(&continueOrQuit)
		if strings.EqualFold(continueOrQuit, "Q") {
			break
		}
	}
}

func getNumberInput() int64 {
	var number int64
	for {
		fmt.Print("Enter a non-negative integer: ")
		fmt.Scanln(&number)
		if number >= 0 {
			break
		}
		fmt.Print("Try again. ")
	}
	return number
}

func getModeInput() Mode {
	var mode Mode
	for {
		fmt.Println("Choose the method of computation.")
		fmt.Print("Enter 1 for recursive or 2 for iterative: ")
		fmt.Scanln(&mode)
		if mode == 1 || mode == 2 {
			break
		}
		fmt.Print("Try again. ")
	}
	return mode
}

func compute(n int64, mode Mode) *big.Int {
	var result *big.Int
	switch mode {
	case Recursive:
		result = factorialRecursive(n)
	case Iterative:
		result = factorialIterative(n)
	}
	return result
}

func factorialRecursive(n int64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	}
	var result big.Int
	result.Mul(big.NewInt(n), factorialRecursive(n-1))
	return &result
}

func factorialIterative(n int64) *big.Int {
	product := big.NewInt(1)
	for i := n; i > 1; i-- {
		product.Mul(product, big.NewInt(i))
	}
	return product
}
