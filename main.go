package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(robMoney(2, 7, 9, 3, 1))
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
