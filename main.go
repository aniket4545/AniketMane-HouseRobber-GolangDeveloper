package main

import (
	"fmt"
	"math"
)

func main() {
	var problemStatement = ` - House Robber Assignment
	Given a list of non-negative integers nums representing the amount of money of each house, 
	return the maximum amount of money you can rob tonight without alerting the police.
	
	Input: %v
	Output: %d
	`
	inputArray := []uint{2, 7, 9, 3, 1}
	output := robMoney(inputArray...)
	fmt.Printf(problemStatement, inputArray, output)
}

func robMoney(houses ...uint) uint {
	return rob(0, make(map[int]uint), houses...)
}

func rob(currentIndex int, amountMap map[int]uint, houses ...uint) uint {

	if currentIndex >= len(houses) {
		return 0
	}

	if amount, exists := amountMap[currentIndex]; exists {
		return amount
	}

	stealedAmount := houses[currentIndex] + rob(currentIndex+2, amountMap, houses...)
	skipHouse := rob(currentIndex+1, amountMap, houses...)

	maxValue := uint(math.Max(float64(stealedAmount), float64(skipHouse)))
	amountMap[currentIndex] = maxValue
	return maxValue
}
