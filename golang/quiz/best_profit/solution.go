package main

import "fmt"

func getBestProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	maxdiff := 0
	for i, min := range prices {
		for j := i; j < len(prices); j++ {
			diff := prices[j] - min
			if diff > maxdiff {
				maxdiff = diff
			}
		}
	}
	return maxdiff
}

func main() {
	prices := []int{6, 0, -1, 10}
	fmt.Println("Prices are ", prices)
	fmt.Println("What is the best profit number?")
	fmt.Println(getBestProfit(prices))
}
