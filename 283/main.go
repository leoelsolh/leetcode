package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

func moveZeroes(nums []int) {
	var count int
	for _, num := range nums {
		if num != 0 {
			nums[count] = num
			count++
		}
	}
	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
}
