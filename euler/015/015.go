/*
Project Euler problem number 15 (https://projecteuler.net/problem=15)

Lattice paths

Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

How many such routes are there through a 20×20 grid?

Solution By Peter Welsh
09/12/2022
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println(getNumPaths(20)) // 137846528820
}

func getNumPaths(gridSize uint64) uint64 {
	numPaths := uint64(1)
	for i := uint64(1); i <= gridSize; i++ {
		numPaths = numPaths * (gridSize + i) / i
	}
	return numPaths
}
