/*
Project Euler problem number 17 (https://projecteuler.net/problem=17)

Number letter counts

If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.

Solution By Peter Welsh
09/12/2022
*/

package main

import (
	"fmt"
)

func main() { 
	length := 0
	for i := 1; i <= 1000; i++ {
		english := getEnglish(i)
		length += len(english)
	}
	fmt.Println(length) // 21124
}

func getEnglish(i int) string {
	if i > 1000 {
		panic("Not implemented.")
	}
	var english string
	i, english = oneThousandPlus(i, english)
	i, english = oneHundredPlus(i, english)
	return oneThruNinetyNine(i, english)
}

func oneThousandPlus(i int, english string) (int, string) {
	if i == 1000 {
		english += getEnglish(i / 1000)
		english += "thousand"
		i = i % 1000
	}
	return i, english
}

func oneHundredPlus(i int, english string) (int, string) {
	if i < 100 {
		return i, english
	}
	english += getEnglish(i / 100)
	english += "hundred"
	if i%100 != 0 {
		english += "and"
	}
	i = i % 100
	return i, english
}

func oneThruNinetyNine(i int, english string) string {
	if i >= 90 {
		english += "ninety"
	} else if i >= 80 {
		english += "eighty"
	} else if i >= 70 {
		english += "seventy"
	} else if i >= 60 {
		english += "sixty"
	} else if i >= 50 {
		english += "fifty"
	} else if i >= 40 {
		english += "forty"
	} else if i >= 30 {
		english += "thirty"
	} else if i >= 20 {
		english += "twenty"
	} else if i >= 10 {
		return tenThruNineteen(i, english)
	}
	i = i % 10
	return oneThruNine(i, english)
}

func tenThruNineteen(i int, english string) string {
	switch i {
	case 19:
		return english + "nineteen"
	case 18:
		return english + "eighteen"
	case 17:
		return english + "seventeen"
	case 16:
		return english + "sixteen"
	case 15:
		return english + "fifteen"
	case 14:
		return english + "fourteen"
	case 13:
		return english + "thirteen"
	case 12:
		return english + "twelve"
	case 11:
		return english + "eleven"
	case 10:
		return english + "ten"
	default:
		return english
	}
}

func oneThruNine(i int, english string) string {
	switch i {
	case 9:
		return english + "nine"
	case 8:
		return english + "eight"
	case 7:
		return english + "seven"
	case 6:
		return english + "six"
	case 5:
		return english + "five"
	case 4:
		return english + "four"
	case 3:
		return english + "three"
	case 2:
		return english + "two"
	case 1:
		return english + "one"
	default:
		return english
	}
}
