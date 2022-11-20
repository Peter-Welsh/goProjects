/*
This program prompts the user for a dollar amount and makes change using standard USD denominations

By Peter Welsh
08-25-2022
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	for {
		amount := getAmountInput()
		makeChange(amount)

		time.Sleep(time.Second)
		var continueOrQuit string
		fmt.Println("\r\nPress enter to restart or Q to quit.")
		fmt.Scanln(&continueOrQuit)
		if strings.EqualFold(continueOrQuit, "Q") {
			break
		}
	}
}

func getAmountInput() float32 {
	var amount float32
	for {
		fmt.Print("Enter dollar amount: ")
		fmt.Scanln(&amount)
		if amount > 0 {
			break
		}
		fmt.Print("Try again. ")
	}
	return amount
}

func makeChange(amount float32) {
	bills, coins := getDenominations()
	amountInBills := int(amount)
	billQuantities := getQuantities(bills, amountInBills)
	printResult(billQuantities, '$', false)

	amountInCents := getAmountInCents(amount)
	coinQuantities := getQuantities(coins, amountInCents)
	printResult(coinQuantities, '¢', true)
}

func getDenominations() (map[int]int, map[int]int) {
	bills := map[int]int{1: 0, 5: 0, 10: 0, 20: 0, 50: 0, 100: 0}
	coins := map[int]int{1: 0, 5: 0, 10: 0, 25: 0}
	return bills, coins
}

func getQuantities(quantities map[int]int, amount int) map[int]int {
	if amount == 0 {
		return nil
	}
	denoms := getKeysDescending(quantities, false)
	for _, denom := range denoms {
		quantity := amount / denom
		if quantity >= 1 {
			quantities[denom] = quantity
			amount -= denom * quantity
		}
	}
	return quantities
}

func getKeysDescending(kvps map[int]int, excludeZeros bool) []int {
	keys := make([]int, 0, len(kvps))
	for key := range kvps {
		if excludeZeros && kvps[key] == 0 {
			continue
		}
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	return keys
}

func getAmountInCents(amount float32) int {
	amountString := fmt.Sprintf("%.2f", amount)
	amountStrings := strings.Split(amountString, ".")
	amountInCents, _ := strconv.Atoi(amountStrings[1])
	return amountInCents
}

func printResult(denominations map[int]int, currencySymbol rune, displaySymbolAfterAmount bool) {
	keys := getKeysDescending(denominations, true)
	if len(keys) == 0 {
		return
	}
	fmt.Println("Denomination\tQty")
	for _, denom := range keys {
		qty := denominations[denom]
		if displaySymbolAfterAmount {
			fmt.Printf("%d%c\t\t%d\r\n", denom, currencySymbol, qty)
		} else {
			fmt.Printf("%c%d\t\t%d\r\n", currencySymbol, denom, qty)
		}
	}
}
