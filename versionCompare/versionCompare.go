/*
This program prompts the user for two versions and compares them.

Examples:
Input: v1 = 1.2.4; v2 = 1.2.3
Output: Version 1 is greater.

Input: v1 = 1.2.3; v2 = 1.3
Output: Version 2 is greater.

Input: v1 = 1.2.0; v2 = 1.2
Output: The versions are the same.

By Peter Welsh
08-23-2022
*/

package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	for {
		version1 := getVersionInput(1)
		version2 := getVersionInput(2)
		compareVersions(version1, version2)

		time.Sleep(time.Second)
		var continueOrQuit string
		fmt.Println("\r\nPress enter to restart or Q to quit.")
		fmt.Scanln(&continueOrQuit)
		if strings.EqualFold(continueOrQuit, "Q") {
			break
		}
	}
}

func getVersionInput(x int) string {
	var version string
	for {
		fmt.Printf("Enter version %v: ", x)
		fmt.Scanln(&version)
		if isValid, _ := regexp.MatchString(`^\d+(\.\d+)*$`, version); isValid {
			break
		}
		fmt.Printf("The version you entered (%v) is invalid. A valid version...\r\n", version)
		fmt.Println("\t1) contains only digits (0-9) or dots (.)")
		fmt.Println("\t2) begins with a digit")
		fmt.Println("\t3) ends with a digit")
		fmt.Println("\t4) contains no more than 1 consecutive dots")
		fmt.Println("Try again.")
	}
	return version
}

func compareVersions(version1 string, version2 string) {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	greaterVersion := getGreaterVersion(v1, v2)
	printResult(greaterVersion)
}

func getGreaterVersion(v1 []string, v2 []string) int {
	for i := 0; float64(i) < math.Max(float64(len(v1)), float64(len(v2))); i++ {
		version1, version2 := convertVersions(i, v1, v2)
		if version1 > version2 {
			return 1
		}
		if version2 > version1 {
			return 2
		}
	}
	return 0
}

func convertVersions(i int, v1 []string, v2 []string) (int, int) {
	var version1 int
	var version2 int
	if i >= len(v1) {
		version1 = 0
	} else {
		version1, _ = strconv.Atoi(v1[i])
	}
	if i >= len(v2) {
		version2 = 0
	} else {
		version2, _ = strconv.Atoi(v2[i])
	}
	return version1, version2
}

func printResult(greaterVersion int) {
	if greaterVersion == 1 || greaterVersion == 2 {
		fmt.Printf("Version %v is greater.\r\n", greaterVersion)
		return
	}
	fmt.Println("The versions are the same.")
}
