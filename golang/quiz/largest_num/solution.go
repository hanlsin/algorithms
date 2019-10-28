package main

import "fmt"

func findLargestNum(numbers []int) int {
	largest := 0
	for _, n := range numbers {
		if (largest < n) {
			largest = n
		}
	}
	return largest
}

func main() {
	numbers := []int{1, 2, 4}
	fmt.Println("Largetst number of ", numbers)
	fmt.Println(findLargestNum(numbers))
}
