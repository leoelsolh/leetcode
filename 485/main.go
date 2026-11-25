package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1}))
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	var best int
	var curr int
	for _, n := range nums {
		if n == 1 {
			curr++
			if curr > best {
				best = curr
			}
		} else {
			curr = 0
		}
	}
	return best
}
