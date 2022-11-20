/*
Project Euler problem number 49 (https://projecteuler.net/problem=49)

Prime permutations

The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways:
(i) each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations of one another.

There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property, but there is one other 4-digit increasing sequence.

What 12-digit number do you form by concatenating the three terms in this sequence?

Solution By Peter Welsh
09/28/2022
*/

package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	const delta = 3330
	const start = 1487 + 1
	answer := getAnswer(start, delta)
	fmt.Println(answer) // 296962999629
}

func getAnswer(start, delta int) int {
	cache := make(map[int]bool)
	var result int
	for i := start; i <= 10000-delta-delta; i++ {
		if !isPrime(i, cache) {
			continue
		}
		next := i + delta
		iStr := strconv.Itoa(i)
		if !isPermutation(iStr, strconv.Itoa(next)) || !isPrime(next, cache) {
			continue
		}
		nextNext := i + delta + delta
		if !isPermutation(iStr, strconv.Itoa(nextNext)) || !isPrime(nextNext, cache) {
			continue
		}
		result, _ = strconv.Atoi(fmt.Sprintf("%d%d%d", i, next, nextNext))
		break
	}
	return result
}

func isPermutation(x, y string) bool {
	return SortString(x) == SortString(y)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func isPrime(n int, cache map[int]bool) bool {
	if isPrime, exists := cache[n]; exists {
		// check if the answer is cached
		return isPrime
	}
	if n <= 1 {
		// all primes are greater than 1
		cache[n] = false
		return false
	}
	if n%2 == 0 || n%3 == 0 {
		// 2 is a prime and 3 is a prime
		// but all other numbers evenly divisible by 2 or 3 are not prime
		cache[n] = n == 2 || n == 3
		return cache[n]
	}
	for i := 5; i <= int(math.Sqrt(float64(n))); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			cache[n] = false
			return false
		}
	}
	cache[n] = true
	return true
}
