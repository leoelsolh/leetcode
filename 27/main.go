package main

import "fmt"

func main() {
	removeElement([]int{3, 2, 2, 3}, 3)
	removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
}

func removeElement(nums []int, val int) int {
	var k int
	for _, n := range nums {
		if n == val {
			continue
		}
		nums[k] = n
		k++
	}
	fmt.Printf("%d, nums = %v\n", k, nums)
	return k
}
