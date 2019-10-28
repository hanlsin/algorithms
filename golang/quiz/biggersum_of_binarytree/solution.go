package main

import "fmt"

func findBiggerSide(numbers []int) string {
	lsum := 0
	rsum := 0
	depth := 0
	idx := 1

	for {
		idx = (2 * depth) + 1
		depthSize := 2 * (depth + 1)
		depth++

		fmt.Println("\nStarting index: ", idx)
		fmt.Println("Depth: ", depth)
		fmt.Printf("Will check all %d of numbers\n", depthSize)

		// calculate left-side
		for i := idx; i < idx + (depthSize/2); i++ {
			if len(numbers) <= i {
				// if no more numbers, break
				break
			}
			fmt.Println("Checking left ... ", numbers[i])
			if numbers[i] < 0 {
				// if less than 0, ignore
				continue
			}
			lsum += numbers[i]
		}

		// calculate right-side
		for i := idx + (depthSize/2); i < idx + depthSize; i++ {
			if len(numbers) <= i {
				// if no more numbers, break
				break
			}
			fmt.Println("Checking right ... ", numbers[i])
			if numbers[i] < 0 {
				// if less than 0, ignore
				continue
			}
			rsum += numbers[i]
		}

		idx += depthSize
		if idx > len(numbers) {
			// checking done.
			break
		}
		fmt.Println("Next starting index: ", idx)
	}

	if lsum > rsum {
		return "Left"
	} else if lsum < rsum {
		return "Right"
	} else {
		return ""
	}
}

func main() {
	numbers := []int{3, 6, 2, 9, -1, 10}
	fmt.Println("Binary Tree numbers: ", numbers)
	ret := findBiggerSide(numbers)
	fmt.Println("\nWho is bigger?")
	fmt.Println(ret)
}
