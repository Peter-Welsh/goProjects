/*
Project Euler problem number 22 (https://projecteuler.net/problem=22)

Names scores

Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order.
Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list.
So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?

Solution By Peter Welsh
09/23/2022
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	score := getAnswer()
	fmt.Println(score) // 871198282
}

func getAnswer() int {
	content, err := os.ReadFile("022/p022_names.txt")
	if err != nil {
		fmt.Println("Err")
	}
	var names []string
	// content is expected to look like: "NAME1", "NAME2", ...
	for _, v := range strings.Split(string(content), ",") {
		name := strings.Trim(v, "\"")
		names = append(names, name)
	}
	sort.Strings(names)
	score := 0
	for i, name := range names {
		value := getValue(name)
		score += value * (i + 1)
	}
	return score
}

func getValue(name string) int {
	value := 0
	for _, v := range name {
		// 'A' has a value of 1, 'B' is 2, etc.
		value += int(v-'A') + 1
	}
	return value
}
