/*
Project Euler problem number 19 (https://projecteuler.net/problem=19)

Counting Sundays

You are given the following information, but you may prefer to do some research for yourself.

    1 Jan 1900 was a Monday.
    Thirty days has September,
    April, June and November.
    All the rest have thirty-one,
    Saving February alone,
    Which has twenty-eight, rain or shine.
    And on leap years, twenty-nine.
    A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

Solution By Peter Welsh
09/14/2022
*/

package main

import (
	"fmt"
)

func main() {
	answer := getNumSundaysBetween(1, 1, 1901, 31, 12, 2000)
	fmt.Println(answer) //171
}

func getNumSundaysBetween(day1, month1, year1, day2, month2, year2 int) int {
	dayOfWeek := 1
	numSundays := 0
	for year := 1900; year <= year2; year++ {
		isLeap := isLeapYear(year)
		for month := 1; month <= 12; month++ {
			daysInMonth := getNumDaysInMonth(month, isLeap)
			for dayOfMonth := 1; dayOfMonth <= daysInMonth; dayOfMonth++ {
				dayOfWeek++
				dayOfWeek = (dayOfWeek % 7) + 1
				qualifies := isSunday(dayOfWeek) && isFirstOfTheMonth(dayOfMonth) && year >= year1
				if qualifies {
					numSundays++
				}
			}
		}
	}
	return numSundays
}

func isFirstOfTheMonth(dayOfMonth int) bool {
	return dayOfMonth == 1
}

func isSunday(dayOfWeek int) bool {
	return dayOfWeek%7 == 1
}

func isLeapYear(year int) bool {
	isDivisibleBy4 := year%4 == 0
	isCentury := year%100 == 0
	is400thCentury := year%400 == 0
	return isDivisibleBy4 && (!isCentury || is400thCentury)
}

func getNumDaysInMonth(i int, isLeapYear bool) int {
	numDays := 0
	switch i {
	case 1: //jan
		numDays = 31
	case 2: //feb
		if isLeapYear {
			numDays = 29
		} else {
			numDays = 28
		}
	case 3: //mar
		numDays = 31
	case 4: //apr
		numDays = 30
	case 5: //may
		numDays = 31
	case 6: //jun
		numDays = 30
	case 7: //jul
		numDays = 31
	case 8: //aug
		numDays = 31
	case 9: //sep
		numDays = 30
	case 10: //oct
		numDays = 31
	case 11: //nov
		numDays = 30
	case 12: //dec
		numDays = 31
	}
	return numDays
}
